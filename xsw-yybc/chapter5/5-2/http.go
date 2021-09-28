package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.New()
	route.POST("", func(ctx *gin.Context) {
		fmt.Println("进来了..")
		b, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg": "出错了.." + err.Error(),
			})
			return
		}
		fmt.Println("body len ==", len(b))
		ctx.JSON(200, gin.H{
			"msg": "你请求的参数是: " + string(b),
		})
	})
	route.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"getMsg": "哈哈哈哈" + time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	route.Run(":8080")
}
