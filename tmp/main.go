package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

type Logger interface {
	LogError(err error)
	GetName() string
}

type Entry struct {
	Name string
}

func (e *Entry) LogError(err error) {
	fmt.Println(err)
}

func (e *Entry) GetName() string {
	return e.Name
}

func main() {
	m := martini.Classic()

	m.Use(func(c martini.Context) {
		e := &Entry{}
		e.Name = "吴彦祖"
		c.MapTo(e, (*Logger)(nil))

		c.Next()
		fmt.Println("11111111111")
		c.Next()
		fmt.Println("222222222222")
		c.Next()
		fmt.Println("33333")

	})

	m.Get("/test", func(req *http.Request, al Logger, w http.ResponseWriter) {
		fmt.Printf("al 类型为 %T , name 值为: %s \n ", al, al.GetName())
		w.Write([]byte("Hello, world!"))
	})
	m.RunOnAddr(":8086")
}
