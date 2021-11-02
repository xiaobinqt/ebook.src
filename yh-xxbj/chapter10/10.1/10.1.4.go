package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())
}
