package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	mid1 := func(c *gin.Context) {
		fmt.Println("mid1 start")
		c.Next()
		fmt.Println("mid1 end")
	}
	mid2 := func(c *gin.Context) {
		fmt.Println("mid2 start")
		//c.Abort()
		c.Next()
		fmt.Println("mid2 end")
	}
	mid3 := func(c *gin.Context) {
		fmt.Println("mid3 start")
		c.Next()
		fmt.Println("mid3 end")
	}
	router.Use(mid1, mid2, mid3)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("process get request")
		c.JSON(http.StatusOK, "hello")
	})
	router.Run()
}
