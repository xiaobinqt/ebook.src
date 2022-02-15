package main

import "fmt"

func luckyNumbers(matrix [][]int) (ans []int) {
	for _, row := range matrix {
		fmt.Println(row)
	next:
		for i, x := range row {
			// 所在行的最小值
			for _, y := range row {
				fmt.Println("开始", x, y)
				if x > y {
					continue next
				}
				fmt.Println("结束======", x, y)
			}

			// 所在列的最大值
			for _, val := range matrix {
				if val[i] > x {
					continue next
				}
			}

			ans = append(ans, x)
		}
	}
	fmt.Println("result ", ans)
	return
}

func main() {
	matrix := [][]int{{10, 1, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	//fmt.Println("origin", matrix)
	//fmt.Println(len(matrix))
	luckyNumbers(matrix)
}
