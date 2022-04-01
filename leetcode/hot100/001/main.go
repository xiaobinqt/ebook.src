package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for i := index + 1; i < len(nums); i++ {
			if (value + nums[i]) == target {
				return []int{index, i}
			}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
