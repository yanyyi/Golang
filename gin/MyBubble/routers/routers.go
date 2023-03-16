package routers

import (
	"MyBubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项

		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
