package routers

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
)

func InitRouters(r *gin.Engine) {
	r.Use(middleware.Others)
	// 登录注册
	InitApi(r)
	// 课程管理
	initCourse(r)

}
