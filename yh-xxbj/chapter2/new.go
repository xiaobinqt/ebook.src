package main

import "fmt"

type T struct {
	ap []string
}

func main() {
	t := new(T)
	t.ap = append(t.ap, "11")
	fmt.Println(t.ap)
}
