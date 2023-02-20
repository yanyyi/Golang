package main

import (
	"github.com/gin-gonic/gin"
	"go_web/routers"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	routers.InitRouters(r)

	r.Run(":8080")
}
