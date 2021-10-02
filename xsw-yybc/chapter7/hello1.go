package main

import "fmt"

func Foo() (int, error) {
	return 0, nil
}

func main() {
	i, _ := Foo()
	fmt.Println(i)
}
