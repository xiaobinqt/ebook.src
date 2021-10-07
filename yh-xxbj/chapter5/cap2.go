package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	s = append(s, 1, 2)
	fmt.Println(s, len(s), cap(s))
}
