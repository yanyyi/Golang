package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型
type User struct {
	gorm.Model   //内嵌gorm。model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` //零值类型,添加指定列名,最大限度保证数据安全
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号(member number)唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string  `gorm:"index:addr"`      //给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段

}

//给结构体在MySQL中重起表名
//func (User) TableName() string {
//	return "myUserTable"  //自动调用,在MySQL中创建一个名为myusertable表
//}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&User{})

	//db.Table("myusertable2").CreateTable(&User{}) //在MySQL中创建一个名为myusertable2表

}
