package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gwp?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, content, author from posts where id > ?"
	var posts []Post
	err := db.Select(&posts, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("posts:%v\n", posts)
}

func main() {
	_ = initDB()
	queryMultiRowDemo()

}
