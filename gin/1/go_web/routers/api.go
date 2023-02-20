package routers

import (
	"github.com/gin-gonic/gin"
	"go_web/web"
)

func InitApi(r *gin.Engine) {
	//http:localhost:8080/api
	apiV1(r)
	apiV2(r)

}

func apiV1(r *gin.Engine) {
	v1 := r.Group("v1")
	api := v1.Group("/api")
	//登录
	api.GET("/login", web.Login)
	//注册
	api.POST("/register", web.Register)
}

func apiV2(r *gin.Engine) {
	v2 := r.Group("v2")
	api := v2.Group("/api")
	//登录
	api.GET("/login", web.Login)
	//注册
	api.POST("/register", web.Register)
}
