package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	fmt.Println("get courses")
}

func AddCourse(c *gin.Context) {
	fmt.Println("add courses")
}

func EditCourse(c *gin.Context) {
	fmt.Println("edit courses")
}

func DelCourse(c *gin.Context) {
	fmt.Println("delete courses")
}
