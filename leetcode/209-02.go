package main

import (
	"fmt"
	"math"
)

func minSubArrayLen02(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	ans := math.MaxInt32
	i := 0
	sum := 0 // 子数组的和

	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func main() {
	fmt.Println(minSubArrayLen02(7, []int{2, 3, 1, 2, 4, 3}))
}
