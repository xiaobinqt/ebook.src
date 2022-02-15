package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.New()

	routeAuth := route.Group("", func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Token")
		if token != "123456" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "需要验证....",
			})
		}
	})

	routeAuth.GET("test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	route.GET("test2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "test2 success",
		})
	})
	route.Run()
}
