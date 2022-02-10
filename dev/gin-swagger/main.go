package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Esign API
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com

//@host 127.0.0.1:1234
func main() {
	route := gin.New()
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger 路由
	route.GET("/test", Test)                                                 // 业务路由
	route.Run(":1234")
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "succ",
	})
}
