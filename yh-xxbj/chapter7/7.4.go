package main

import "fmt"

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string {
		fmt.Println("11111111")
		return "hello world!"
	})

	fmt.Println(t)
}
