package main

import "fmt"

func isValid(s string) bool {
	stack := make([]string, 0)
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	if s == "" {
		return false
	}

	sRune := []rune(s)
	if len(sRune)%2 != 0 {
		return false
	}

	for _, each := range sRune {
		str := string(each)
		if str == "(" || str == "{" || str == "[" {
			stack = append(stack, str)
			continue
		}

		if v, ok := m[str]; ok {
			if len(stack) > 0 && v == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	fmt.Println(stack)
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("]]"))
}
