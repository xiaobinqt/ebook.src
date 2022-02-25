package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.New()
	route.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "111")
	})

	route.Run()
}
