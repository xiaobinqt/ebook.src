package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int

	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]

		// if n1 > target { // 不能这样写,因为可能是负数
		// 	break
		// }

		if i > 0 && n1 == nums[i-1] { // 对 nums[i] 去重
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 对 nums[j] 去重
				continue
			}

			left := j + 1
			right := len(nums) - 1
			for left < right {
				n3 := nums[left]
				n4 := nums[right]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for left < right && n3 == nums[left+1] { // 去重
						left++
					}
					for left < right && n4 == nums[right-1] { // 去重
						right--
					}

					// 找到答案时,双指针同时靠近
					right--
					left++
				}
			}
		}
	}

	return res
}

func main() {
	ret := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(ret)
}
