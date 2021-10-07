package main

import "fmt"

func main() {
	s := make([]int, 0, 100)
	s1 := s[:2:4]
	fmt.Println(s1, cap(s1))

}
