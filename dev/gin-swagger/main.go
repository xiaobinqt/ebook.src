package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go.src/dev/gin-swagger/docs"
)

// 参考文档 https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html#mime-types

// @title 我的 gin swagger API
// @version 1.0
// @description  我的 gin swagger API 描述
// @securityDefinitions.apiKey ApiToken
// @in header
// @name X-Token
func main() {
	route := gin.New()

	server := &http.Server{
		Addr:    ":1234",
		Handler: route,
	}

	authRoute := route.Group("", func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Token")
		if token != "123456" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "no token unauthorized",
			})
		}
	})

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger 路由
	// header X-Token 需要验证
	authRoute.GET("/api/get/:id", Get)
	route.POST("/api/post/:id", Post)
	route.POST("/api/post/formdata", FormData)

	server.ListenAndServe()
}

// @Tags 组1
// @Summary 一个 get 请求
// @Security ApiToken
// @Param id path string true "ID 信息"
// @Param page query int false "第几页数据"
// @Param size query int false "每页多少条数据"
// @Success 200 {object} ReturnMessage
// @Router /api/get/{id} [GET]
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ReturnMessage{
		Msg: fmt.Sprintf("success, 请求 id 为: %s, page:%s, size: %s",
			ctx.Param("id"), ctx.Query("page"), ctx.Query("size")),
		Code: 0,
	})
}

// @Tags 组2
// @Summary 一个 post 请求
// @Param id path string true "ID"
// @Param body body	PostReq true "JSON数据"
// @Success 200 {object} ReturnMessage
// @Router /api/post/{id} [POST]
func Post(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ReturnMessage{
		Msg:  "success",
		Code: 0,
	})
}

// @Tags 组2
// @Summary 需要 formdata 参数
// @Param name formData string true "姓名"
// @Param upload_file formData file true "需要上传的文件"
// @Success 200 {object} ReturnMessage
// @Router /api/post/formdata [POST]
func FormData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ReturnMessage{
		Msg:  "success",
		Code: 0,
	})
}

type PostReq struct {
	Name string `json:"name"`
}

type ReturnMessage struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
