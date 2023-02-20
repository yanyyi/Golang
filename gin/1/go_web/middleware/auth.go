package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const token = "123456"

func AuthCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_token")
	fmt.Println("access token is :" + accessToken)
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token 校验失败",
		})
		c.AbortWithError(http.StatusInternalServerError, errors.New("token check failed"))
	}
	c.Set("user_name", "nick")
	c.Set("user_id", "10001")
	fmt.Println("auth Check end")
	c.Next()
}
