package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums) // 先排序

	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	ans := math.MaxInt64
	for i, num := range nums[:len(nums)-k+1] { // 左闭右开区间
		ans = min(ans, nums[i+k-1]-num)
	}

	return ans
}

func main() {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}
