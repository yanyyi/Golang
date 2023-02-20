package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// ---------------------------------------------增------------------------------------------//
	//创建表 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{1, "姚明", "男", "打篮球"}
	u2 := UserInfo{2, "梅西", "男", "踢足球"}
	db.Create(&u1)
	db.Create(&u2)

	// ---------------------------------------------查------------------------------------------//
	var u UserInfo
	db.First(&u) //查询表中第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)

	// ---------------------------------------------改------------------------------------------//
	db.Model(&u).Update("hobby", "退休后当篮球教练") //可在数据库查看

	// ---------------------------------------------删------------------------------------------//
	db.Delete(&u) //可在数据库查看
}
