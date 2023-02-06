package main

import "fmt"

func erfen(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		middle := (left + right) / 2
		if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	fmt.Println(erfen([]int{1, 2, 5, 8, 10, 18}, 18))
	fmt.Println(erfen([]int{1}, 1))
}
