package main

import "fmt"

func main() {
	data := [3]string{"a", "b", "c"}

	s := data[:]
	s = append(s, "1")
	fmt.Println(s)
	fmt.Println(data)
}
