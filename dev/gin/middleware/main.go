package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.New()

	route.Use(func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Token")
		if token != "123456" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "token invalid unauthorized",
			})
		}
	}).GET("test1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "success")
	})

	route.Run()
}
