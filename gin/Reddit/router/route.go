package router

import (
	"Reddit/controller"
	"Reddit/middleware"

	//"Reddit/controller"
	"Reddit/logger"
	//"Reddit/middleware"
	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	//
	//注册
	//v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/signup", controller.SignUpHandler)
	////登录
	v1.POST("/login", controller.LoginHandler)
	//
	v1.Use(middleware.JWTAuthMiddleware()) //应用JWT认证中间件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}
	////r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	////	//如果是登录的用户,判断请求头中是否有有效的 JWT?
	////	c.String(http.StatusOK, "pong")
	////})
	//{
	//	v1.GET("/community", controller.CommunityHandler)
	//	v1.GET("/community/:id", controller.CommunityDetailHandler)
	//
	//	v1.POST("/post", controller.CreatePostHandler)
	//	v1.GET("/post/:id", controller.GetPostDetailHandler)
	//	v1.GET("/posts", controller.GetPostListHandler)
	//
	//	//根据时间或分数获取帖子列表
	//	v1.GET("/posts2", controller.GetPostListHandler2)
	//
	//	//投票
	//	v1.POST("/vote", controller.PostVoteController)
	//}
	//
	//pprof.Register(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
