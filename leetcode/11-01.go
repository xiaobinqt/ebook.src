package main

import "fmt"

/**
我艹，暴力法提交`超出时间限制`
*/
func maxArea(height []int) int {
	max := 0
	getMax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	getMix := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * getMix(height[i], height[j])
			max = getMax(area, max)
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
