package main

import "fmt"

func maxArea02(height []int) int {
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

	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := getMix(height[left], height[right]) * (right - left)
		max = getMax(area, max)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea02([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
