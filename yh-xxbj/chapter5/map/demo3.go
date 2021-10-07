package main

import "fmt"

func test(x map[string]int) {
	x["x"] = 11
}

func main() {
	m := make(map[string]int)
	test(m)

	fmt.Println(m)
}
