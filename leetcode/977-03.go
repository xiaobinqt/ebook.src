package main

import (
	"fmt"
	"strconv"
	"time"
)

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var (
		result = make([]int, len(nums))
		i, j   = 0, len(nums) - 1
		k      = len(nums) - 1 // 能访问到最后一个值
	)

	for i <= j {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
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

type ListNodex struct {
	Val  int
	Next *ListNodex
}

func main() {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))

}
