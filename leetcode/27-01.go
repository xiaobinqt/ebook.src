package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	low := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[low] = nums[i]
			low++
		}
	}

	return low
}

func main() {
	x := []int{1, 2, 6, 7, 9, 6, 6, 6}
	xx := removeElement(x, 6)
	fmt.Println(x, xx)
}
