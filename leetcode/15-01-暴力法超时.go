package main

import (
	"fmt"
	"sort"
)

// 暴力法
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) <= 2 || len(nums) == 3 && nums[0]+nums[1]+nums[2] != 0 {
		return result
	}

	sort.Ints(nums)
	m := make(map[string]interface{}, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					t := []int{nums[i], nums[j], nums[k]}
					sort.Ints(t)
					key := fmt.Sprintf("%d%d%d", t[0], t[1], t[2])
					if _, ok := m[key]; ok == false {
						m[key] = nil
						result = append(result, t)
					}
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
