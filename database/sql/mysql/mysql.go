package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bubble")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// 开启一个事务
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	_, err = tx.Exec("insert into todos (title, status) values (?, ?)", "旅游", "0")
	if err != nil {
		// 如果错误,回滚
		tx.Rollback()
		panic(err.Error())
	}
	_, err = tx.Exec("update todos set status = ? where title=?", "1", "旅游")
	if err != nil {
		// 如果错误,回滚
		tx.Rollback()
		panic(err.Error())
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
}
