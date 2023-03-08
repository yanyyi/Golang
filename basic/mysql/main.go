package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
	age      int
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//插入数据
	var name string
	var age int
	fmt.Println("请输入用户的姓名:")
	fmt.Scan(&name)
	fmt.Println("请输入用户的年龄:")
	fmt.Scan(&age)
	sqlStr := "insert into users (username, age) values (?, ?)"
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, error:%v\n", err)
		return
	}
	lastId, _ := ret.LastInsertId()
	fmt.Println("insert success, last insert id:", lastId)

}
