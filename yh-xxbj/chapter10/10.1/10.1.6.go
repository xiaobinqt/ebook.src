package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	var s http.Server

	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

}
