package main

import (
	"MyBubble/dao"
	"MyBubble/models"
	"MyBubble/routers"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close() //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
