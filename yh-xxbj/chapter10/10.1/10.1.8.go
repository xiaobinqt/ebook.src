package main

import (
	"fmt"
	"reflect"
)

type X2 struct {
}

func (X2) String() string {
	return ""
}

func main() {
	var a X2
	t := reflect.TypeOf(a)

	// Implements 不能直接使用类型作为参数，导致这种用法非常别扭
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))

	it := reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it))
	fmt.Println(t.AssignableTo(st), t.AssignableTo(it))
}
