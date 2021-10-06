package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"go.ebook.src/dhsjms/chapter1/cal_v2/cal"
)

func main() {
	var (
		operate, a, b      string
		aFloat64, bFloat64 float64
		err                error
	)

	read := bufio.NewReader(os.Stdin)
	os.Stdout.WriteString("暂时只支持整数\n")
	os.Stdout.WriteString("请输入数字A: ")
	a, _ = read.ReadString('\n') // 有 \r\n\
	a = strings.Replace(a, "\r\n", "", -1)

	aFloat64, err = strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalf("a err: %s \n", err.Error())
	}

	os.Stdout.WriteString("请输入数字运算符(+,-.*./): ")
	operate, _ = read.ReadString('\n')
	operate = strings.Replace(operate, "\r\n", "", -1)

	if operate != "+" && operate != "-" && operate != "*" && operate != "/" {
		log.Fatalf("operate err %s \n", err.Error())
	}

	os.Stdout.WriteString("请输入数字B: ")
	b, _ = read.ReadString('\n')
	b = strings.Replace(b, "\r\n", "", -1)

	bFloat64, err = strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatalf("b err: %s \n", err.Error())
	}

	var oper cal.OperationInterface
	oper = (&cal.OperationFactory{}).CreateOperate(operate)
	oper.SetNumberA(aFloat64)
	oper.SetNumberB(bFloat64)
	ret := oper.GetResult()
	fmt.Println("最后的结果为 =", ret)
}
