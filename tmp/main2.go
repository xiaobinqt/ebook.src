package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

type Stu struct {
	Name string
}

func Test(s1 *Stu, w http.ResponseWriter) {
	fmt.Println(s1.Name)
	var s2 Stu
	s2.Name = "李白"
	w.Write([]byte("hello world" + s1.Name + "  " + s2.Name))
}

func main() {
	m := martini.Classic()
	var s1 Stu
	s1.Name = "吴彦祖"
	m.Map(&s1)

	m.Group("/templates/:template_id", func(r martini.Router) {
		r.Get("", Test)
		r.Get("/revisions/:revision_id", Test)
	}, func(params martini.Params, req *http.Request) {
		fmt.Println(req.URL.Path)
		fmt.Println(params["template_id"], params["revision_id"])

	})

	m.Get("/test", Test)
	m.RunOnAddr(":8086")
}
