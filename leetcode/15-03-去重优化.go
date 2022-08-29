package main

import (
	"fmt"
	"sort"
)

/**
参考
https://leetcode.cn/problems/3sum/solution/suan-fa-si-wei-yang-cheng-ji-er-fen-cha-5bk43/

三数相加，在左右指针中，如果两个数定了，那么最后一个数也定了，所以当 target 固定时，只要有一个数出现过，那么就可以直接跳过
*/

func threeSum03(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var res = make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		var target = -nums[i]
		var left, right = i + 1, len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right {
					left++
					if nums[left-1] != nums[left] {
						break
					}
				}

				for left < right {
					right--
					if nums[right] != nums[right+1] {
						break
					}
				}

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	// -4 -1 -1 0 1 2
	fmt.Println(threeSum03([]int{-1, 0, 1, 2, -1, -4}))
}
