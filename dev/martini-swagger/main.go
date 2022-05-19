package main

import (
	"embed"
	"net/http"

	"github.com/go-martini/martini"
)

//go:embed swagger
var embededFiles embed.FS

// @title 测试 API
// @version 4.0
// @description 测试 API V4.0
// @securityDefinitions.apiKey MyApiKey
// @in header
// @name Xiaobinqt-Api-Key
// @BasePath /
func main() {
	m := martini.Classic()
	m.Get("/swagger/**", http.FileServer(http.FS(embededFiles)).ServeHTTP)
	m.Post("/api/login/:user_id", Login)
	m.Run()
}

type Req struct {
	Email    string `json:"email"` // 邮箱
	Password string `json:"password"`
}

// @Tags 用户管理
// @Summary 用户登录
// @Security MyApiKey
// @accept application/json
// @Produce application/json
// @Param  user_id path string true "用户ID"
// @Param  search query string false "搜索内容"
// @Param data body Req true "email: 用户名，password: 密码"
// @Success 200 {object} EdgeInstanceList
// @Router /api/login/{user_id} [POST]
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

type EdgeInstanceList struct {
	A string
	B string
}
