package main

import "fmt"

func main() {
	str := "此夜曲中闻折柳"
	strrune := []rune(str)
	fmt.Println(strrune)
	for i, x := range str {
		fmt.Println(i, string(x))
	}
}
