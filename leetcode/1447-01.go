package main

import (
	"fmt"
	"strconv"
)

// 可以利用欧几里得算法，参考 https://www.cnblogs.com/csmSimona/p/12011582.html
func simplifiedFractions(n int) (ans []string) {
	var f = func(a, b int) int { // a > b
		// 2个数相除，得出余数
		// 如果余数不为0，则拿较小的数与余数继续相除，判断新的余数是否为0
		// 如果余数为0，则最大公约数就是本次相除中较小的数。
		// 比如数字 25 和 10 ，使用辗转相除法求最大公约数过程如下：
		// 25 除以 10 商 2 余 5
		// 根据辗转相除法可以得出，25 和 10 的最大公约数等于 5 和 10 之间的最大公约数
		// 10 除以 5 商 2 余 0， 所以 5 和 10 之间的最大公约数为 5，因此25 和 10 的最大公约数为 5
		for b != 0 { // 分母不能为0
			b, a = a%b, b
		}
		return a
	}
	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if f(denominator, numerator) == 1 {
				ans = append(ans, strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator))
			}
		}
	}
	return
}

func main() {
	//var f = func(a, b int) int {
	//	for b != 0 { // 分母不能为0
	//		b, a = a%b, b
	//		fmt.Println(b, a)
	//	}
	//	return a
	//}
	//
	//fmt.Println(f(25, 10))
	fmt.Println(simplifiedFractions(10))
}
