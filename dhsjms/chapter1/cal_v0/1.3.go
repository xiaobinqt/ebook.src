package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		operator, a, b, result string
	)
	read := bufio.NewReader(os.Stdin)
	os.Stdout.WriteString("暂时只支持整数\n")
	os.Stdout.WriteString("请输入数字A: ")
	a, _ = read.ReadString('\n')

	os.Stdout.WriteString("请输入数字运算符(+,-.*./): ")
	operator, _ = read.ReadString('\n')

	os.Stdout.WriteString("请输入数字B: ")
	b, _ = read.ReadString('\n')

	switch operator {
	case "+":

	case "-":
	case "*":
	case "/":

	}
	fmt.Println(operator, a, b, result)
	fmt.Println(a)
	fmt.Println(b)
}
