package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	//1.连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	u := User{Name: "qimi", Age: 18} //在代码层面创建一个User对象
	//fmt.Println(db.NewRecord(&u))    //判断主键是否为空,相当于查看表中是否有数据 此处返回true
	db.Create(&u)
	//fmt.Println(db.NewRecord(&u)) //  此处返回false
}
