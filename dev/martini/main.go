package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(requestCtxMidware)
	m.Get("/test", AH(CateUser, ActCreate), Test)
	m.Run()
}

func requestCtxMidware(c martini.Context, w http.ResponseWriter, r *http.Request) {
	ctx := &RequestContext{
		req:    r,
		res:    w,
		params: make(map[string]string),
		mc:     c,
	}

	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			ctx.params[k] = v[0]
		}
	}

	fmt.Printf("Map 前 Svr 类型为: %T \n", c)

	c.Map(ctx)

	fmt.Printf("Map 前 Svr 类型为: %T \n", c)

	c.Next()
}

func Test(ctx *RequestContext, al Logger) {
	ctx.LogOldVal = "11111"
	al.LogResources(R{RKeyUser: {"xxxxxx"}})
	ctx.res.Write([]byte("test"))
}

// Context for each request
type RequestContext struct {
	req       *http.Request
	res       http.ResponseWriter
	params    martini.Params
	mc        martini.Context
	db        *mgo.Database
	internal  bool
	LogOldVal interface{}
	LogNewVal interface{}
}

func AH(category, action string) martini.Handler {
	return func(ctx *RequestContext, w http.ResponseWriter, req *http.Request, c martini.Context) {
		LogHandler(category, action, "q11@qq.com", w, req, c, ctx)
	}
}
