package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {

	return 0

}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val { // 发现需要移除的元素，就将数组集体向前移动一位
			for j := i + 1; j < length; j++ {
				nums[j-1] = nums[j]
			}

			i--      // 因为下标i以后的数值都向前移动了一位，所以i也向前移动一位
			length-- // 此时数组的大小-1
		}

	}

	return length
}

func main() {
	x := []int{1, 2, 6, 7, 9, 6, 6, 6}
	xx := removeElement(x, 6)
	fmt.Println(x, xx)
}
