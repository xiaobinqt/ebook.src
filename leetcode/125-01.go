package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	str2 := make([]rune, 0)
	for _, each := range []rune(str) {
		if unicode.IsLetter(each) || unicode.IsDigit(each) {
			str2 = append(str2, each)
		}
	}

	reverseString := func(s string) string {
		runes := []rune(s)
		for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
			runes[from], runes[to] = runes[to], runes[from]
		}
		return string(runes)
	}

	fmt.Println(string(str2))
	fmt.Println(reverseString(string(str2)))
	return string(str2) == reverseString(string(str2))
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
