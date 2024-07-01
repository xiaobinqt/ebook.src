package main

import "github.com/gin-gonic/gin"

/**
二叉树最大数字所在的深度,

输入一个链表，求解链表中环的长度

*/

func main() {
	route := gin.New()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"clientIP": ctx.ClientIP(),
			"remoteIP": ctx.RemoteIP(),
		})
	})
	route.Run(":8022")
}
