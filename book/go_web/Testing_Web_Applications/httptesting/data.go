package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
	"time"
)

var db *sql.DB

// 连接到数据库
func initSQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gwp"
	//去初始化全局的db对象,而不是新声明一个db变量
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
		return
	}
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200) //最大连接数
	db.SetMaxIdleConns(10)  //最大空闲连接数
	return
}

// 获取单条指定的帖子
func SelectAPost(id int) (post Post, err error) {
	sqlStr := "select id, content, author from posts where id=?"

	err = db.QueryRow(sqlStr, id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		fmt.Println("找不到指定的帖子")
		return
	}
	return
}

// 获取多条指定的帖子
func SelectPosts() (posts []Post, err error) {
	sqlStr := "select id, content, author from posts limit ?"
	rows, err := db.Query(sqlStr, 8)
	if err != nil {
		fmt.Printf("query failed, error:%v\n", err)
		return
	}
	defer rows.Close()

	//循环读取结果中的数据
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			fmt.Printf("scan failed, error:%v\n", err)
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 创建一篇新帖子
func (post *Post) Create() (err error) {
	sqlStr := `insert into posts (content, author) values (?, ?)`
	ret, err := db.Exec(sqlStr, post.Content, post.Author)
	if err != nil {
		fmt.Printf("insert failed, error:%v\n", err)
		return
	}
	var theId int64
	theId, err = ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, error:%v\n", err)
		return
	}
	fmt.Printf("create post success, the id is %d.\n", theId)
	return
}

// 删除指定的帖子
func (post *Post) Delete() (err error) {
	sqlStr := `delete from posts where id=?`
	_, err = db.Exec(sqlStr, post.Id)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	return
}

// 更新指定的帖子
func (post *Post) Update() (err error) {
	sqlStr := `update posts set content=?, author=? where id=?`
	_, err = db.Exec(sqlStr, post.Content, post.Author, post.Id)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	return
}

//func main() {
//	if err := initSQL(); err != nil {
//		fmt.Printf("connect to db failed, err:%v\n", err)
//	}
//	defer db.Close()
//	fmt.Println("connect to db success!")
//	//p1 := Post{
//	//	Id:      3,
//	//	Content: "Mxnet",
//	//	Author:  "Amazon",
//	//}
//	//p1.Create()
//	//p1.Update()
//	//p1.Delete()
//	llist1, _ := SelectAPost(3)
//	//llist2, _ := SelectPosts()
//	fmt.Println(llist1)
//}
