package main

import "fmt"

func A(m map[string]string) {
	m["1"] = "22222"
}
func main() {
	var m = make(map[string]string)
	m["1"] = "00000"
	fmt.Println(m)

	A(m)
	fmt.Println(m)

}
