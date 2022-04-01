package main

import "fmt"

func isValid(s string) bool {
	n := len(s)

	if n%2 != 0 { // 奇数直接退出
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 如果是右括号,判断栈顶是否是对应的左括号,果然有对应的左括号,则弹出栈顶元素,否则直接退出
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
