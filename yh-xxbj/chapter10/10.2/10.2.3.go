package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		Name string
		Age  int
	}

	u := user{
		Name: "q.yuhen",
		Age:  60,
	}

	v := reflect.ValueOf(&u)

	if v.CanInterface() == false {
		println("CanInterface:fail")
		return
	}

	p, ok := v.Interface().(*user)
	if ok == false {
		println("Interface: fail")
		return
	}

	p.Age++
	fmt.Printf("%+v \n", u)
}
