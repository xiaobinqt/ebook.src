package main

import (
	"fmt"
	"unsafe"

	"ebook.src/yh-xxbj/chapter9/lib"
)

func main() {
	d := lib.NewData()
	d.Y = 200 // 直接访问导出字段

	// 利用指针转换访问私有字段
	p := (*struct {
		x int
	})(unsafe.Pointer(d))
	p.x = 100
	fmt.Printf("%+v \n", *d)
}
