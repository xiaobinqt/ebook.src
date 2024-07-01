package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target {
				ans = min(ans, j-i+1)
				break
			}
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
