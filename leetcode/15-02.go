package main

import (
	"fmt"
	"sort"
)

/**
参考
https://leetcode.cn/problems/3sum/solution/suan-fa-si-wei-yang-cheng-ji-er-fen-cha-5bk43/
*/

func threeSum02(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	m := make(map[string]interface{}, 0)

	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				key := fmt.Sprintf("%d%d%d", nums[i], nums[left], nums[right])
				if _, ok := m[key]; !ok {
					ret = append(ret, []int{nums[i], nums[left], nums[right]})
					m[key] = nil
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ret
}

func main() {
	// -4 -1 -1 0 1 2
	fmt.Println(threeSum02([]int{-1, 0, 1, 2, -1, -4}))
}
