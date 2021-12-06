package main

import "fmt"

func main() {
	var a, b []byte
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	a = ar[2:5]
	b = ar[3:5]

	fmt.Println(a, b)

	a[1] = 'o'
	fmt.Println(a, b)
}
