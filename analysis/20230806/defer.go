package main

import "fmt"

func main() {
	//defer func() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("1111", e)
		}
	}()
	//}()

	panic(404)
}
