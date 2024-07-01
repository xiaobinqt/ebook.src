package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/event", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		closeNotify := c.Writer.(http.CloseNotifier).CloseNotify()
		ticker := time.NewTicker(5 * time.Second)

		for {
			select {
			case <-closeNotify:
				// 当连接关闭时，停止发送事件并退出循环
				fmt.Println("Connection closed")
				ticker.Stop()
				return
			case <-ticker.C:
				// 定时发送事件
				message := "Event message"
				data := fmt.Sprintf("data: %s\n\n", message)
				c.Writer.Write([]byte(data))
				c.Writer.Flush()
			}
		}
	})

	r.Run(":8000")
}
