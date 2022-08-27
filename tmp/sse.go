package main

import (
	"fmt"
	"net/http"
	"time"
)

type SSE struct {
}

func (sse *SSE) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	for {

		select {
		case <-req.Context().Done():
			fmt.Println("req done...")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Fprintf(rw, "id: %d\nevent: ping \ndata: %d\n\n", time.Now().Unix(), time.Now().Unix())
			flusher.Flush()
		}

	}

}

func main() {
	//route := gin.New()
	//route.GET("sse", gin.WrapH(&SSE{}))
	//route.Run(":8080")

	http.Handle("/sse", &SSE{})
	http.ListenAndServe(":8080", nil)
}
