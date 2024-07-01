package main

import "fmt"

func f() int {
	t := 5

	defer func() {
		t = t + 5
	}()

	return t

}

func main() {
	fmt.Println(f())
}
