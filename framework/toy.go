package main

import (
	"net/http"
	"toy"
)

type UserData struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	r := toy.Default()

	r.POST("/user", func(c *toy.Context) {
		var userData UserData

		// 使用 ShouldBindJSON 将请求中的 JSON 数据绑定到结构体
		if err := c.ShouldBindJSON(&userData); err != nil {
			c.JSON(http.StatusBadRequest, toy.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, toy.H{
			"message":  "User data received and processed",
			"username": userData.Username,
			"age":      userData.Age,
		})
	})

	r.Run(":8080")
}
