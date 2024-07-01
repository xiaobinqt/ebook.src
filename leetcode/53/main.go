package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	result := math.MinInt64
	length := len(nums)

	for i := 0; i < length; i++ {
		count = 0
		for j := i; j < length; j++ {
			count += nums[j]
			if count > result {
				result = count
			}
		}
	}

	return result
}

func main() {

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
