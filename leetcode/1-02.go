package main

import "fmt"

func twoSum02(nums []int, target int) []int {
	var m = make(map[int]int, 0)
	for index, each := range nums {
		m[each] = index
	}

	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if idx, ok := m[key]; ok {
			if idx == i {
				continue
			}
			return []int{i, idx}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum02([]int{3, 2, 4}, 6))
}
