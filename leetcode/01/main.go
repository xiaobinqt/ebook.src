package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var (
		m = make(map[int]int, len(nums))
	)

	for index, each := range nums {
		m[each] = index
	}

	for index, val := range m {
		sub := target - val
		if key, ok := m[sub]; ok {
			if key != index {
				return []int{index, sub}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
