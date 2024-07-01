package main

import "fmt"

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	/**
	for循环中变量定义成i或j的细节：按照通常的思维，i代表行，j代表列
	这样，就可以很容易区分出来变化的量应该放在[][]的第一个还是第二个
	  对于变量的边界怎么定义：
	  	从左向右填充：填充的列肯定在[left,right]区间
	  	从上向下填充：填充的行肯定在[top,bottom]区间
	  	从右向左填充：填充的列肯定在[right,left]区间
	  	从下向上填充：填充的行肯定在[bottom,top]区间
	  通过上面的总结会发现边界的起始和结束与方向是对应的
	*/

	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
