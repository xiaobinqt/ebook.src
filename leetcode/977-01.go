package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	for index, item := range nums {
		nums[index] = item * item
	}

	sort.Ints(nums)
	return nums
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
