package main

import "fmt"

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var (
		result = make([]int, len(nums))
		i, j   = 0, len(nums) - 1
		k      = len(nums) - 1
	)

	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			result[k] = nums[j] * nums[j]
			k--
			j--
		}
	}

	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
