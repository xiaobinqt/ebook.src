package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 100

	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)

	fmt.Println(tx, tp, tx == tp)
	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx.Kind(), tp.Elem())
	fmt.Println(tx == tp.Elem())
}
