package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		print(err)
	}
	defer db.Close()

	//创建表  自动迁移(把结构体和数据表对应)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{ID: 1, Name: "姚明", Gender: "男", Hobby: "蛙泳"}
	db.Create(&u1)

	//查询数据行
	var u UserInfo
	db.First(&u) //查询表中第一天数据保存到u中
	fmt.Printf("u:%v\n", u)

	//更新
	db.Model(&u).Update("hobby", "篮球")
	fmt.Printf("u:%v\n", u)

	//删除
	db.Delete(&u)

}
