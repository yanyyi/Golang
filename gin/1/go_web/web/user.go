package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	//userName,_ := c.GetQuery("user_name")
	userName := c.DefaultQuery("user_name", "")
	pwd := c.DefaultQuery("pwd", "")
	c.JSON(http.StatusOK, gin.H{
		"user_name": userName,
		"pwd":       pwd,
	})

	c.JSON(http.StatusOK, gin.H{
		"messgae": "login user",
	})
}

type RegisterIn struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Pwd      string `form:"pwd" json:"pwd" binding:"required"`
	Phone    string `form:"phone" json:"phone" `
	Email    string `form:"email" json:"email" binding:"email"`
}

func Register(c *gin.Context) {
	in := new(RegisterIn)
	err := c.Bind(in)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, in)
}
