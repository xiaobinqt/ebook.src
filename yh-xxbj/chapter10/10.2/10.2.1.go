package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100

	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
}
