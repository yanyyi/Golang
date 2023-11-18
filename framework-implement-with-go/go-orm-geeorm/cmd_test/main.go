package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	geeorm "go-orm-geeorm"
)

func main() {
	engine, _ := geeorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/geeorm")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)

}
