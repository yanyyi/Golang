package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Others(c *gin.Context) {
	fmt.Println("called middleware Others... start")
	c.Next()
	fmt.Println("called middleware Others... end")
}
