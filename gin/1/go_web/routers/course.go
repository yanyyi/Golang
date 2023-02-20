package routers

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/web"
)

func initCourse(r *gin.Engine) {
	r.Use(middleware.AuthCheck)
	courseV1(r)
	courseV2(r)
}

func courseV1(r *gin.Engine) {
	v := r.Group("/v1")
	v.GET("/course", web.GetCourse)
	v.POST("/course", web.AddCourse)
	v.PUT("/course", web.EditCourse)
	v.DELETE("/course", web.DelCourse)
}

func courseV2(r *gin.Engine) {
	v := r.Group("/v2")
	v.GET("/course", web.GetCourse)
	v.POST("/course", web.AddCourse)
	v.PUT("/course", web.EditCourse)
	v.DELETE("/course", web.DelCourse)
}
