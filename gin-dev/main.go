package main

import (
	"fmt"
	"net/http"

	"gin-dev/pkg/setting"
	"gin-dev/routers"
)

// @title go-dev API
// @version 1.0
// @description An example of gin
// @license.name MIT
// @host 127.0.0.1:8000
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
