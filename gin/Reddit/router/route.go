package router

import (
	"Reddit/controller"
	"Reddit/logger"
	"Reddit/middleware"
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

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts", controller.GetPostListHandler)

	v1.Use(middleware.JWTAuthMiddleware()) //应用JWT认证中间件
	{
		v1.POST("/post", controller.CreatePostHandler)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
