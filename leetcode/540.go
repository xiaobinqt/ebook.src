package main

import "fmt"

/**
https://leetcode-cn.com/problems/single-element-in-a-sorted-array/solution/pythonjavajavascriptgo-er-fen-cha-zhao-b-l76s/
https://leetcode-cn.com/problems/single-element-in-a-sorted-array/solution/you-xu-shu-zu-zhong-de-dan-yi-yuan-su-by-y8gh/
**/

// 1, 1, 2, 3, 3, 4, 4, 8, 8
func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 1
			} else {
				high = mid
			}
			continue
		}

		if nums[mid] == nums[mid-1] {
			low = mid + 1
		} else {
			high = mid

		}
	}

	return nums[low]
}

func singleNonDuplicateV0(nums []int) int {
	var m = make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	for v, count := range m {
		if count == 1 {
			return v
		}
	}

	return -1
}

func main() {
	//fmt.Println(3, 3^1)
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
