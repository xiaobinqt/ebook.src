package main

import "fmt"

func main() {
	src := "你好abc啊哈哈"
	dst := reverse([]rune(src))
	fmt.Printf("%v\n", string(dst))
}

func reverse(s []rune) []rune {
	if len(s) == 0 {
		return []rune{}
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
