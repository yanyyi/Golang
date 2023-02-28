package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
	"time"
)

var db *sql.DB

func initSQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1"
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

type user struct {
	id   int
	age  int
	name string
}

// 查询单条示例
func queryOneRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	//非常重要,确保QueryRow之后调用Scan方法,否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("Scan failed, error:%v\n", err)
		return
	}
	fmt.Printf("id:%d, name:%s, age:%d\n", u.id, u.name, u.age)
}

// 查询多条示例
func queryMultiRow() {
	sqlStr := "select id, name, age from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, error:%v\n", err)
		return
	}
	defer rows.Close()

	//循环读取结果中的数据
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, error:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入、更新和删除操作都使用Exec方法
func insertRow(name string, age int) {
	sqlStr := "insert into user (name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, name, age)
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
	fmt.Printf("insert success, the id is %d.\n", theId)
}

// 更新数据
func updateRow(age int, id int) {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, error:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() //影响操作的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, error:%v\n", err)
		return
	}
	fmt.Printf("update success, affectd rows:%d\n", n)
}

func main() {
	if err := initSQL(); err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}
	defer db.Close()
	fmt.Println("connect to db success!")
	//queryOneRow()
	//insertRow("老默", 38)
	updateRow(39, 11) //这里有sql注入的问题，绝对不能让用户去拼接sql语句
	queryMultiRow()
}
