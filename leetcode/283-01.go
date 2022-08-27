package main

import "fmt"

func MoveTarget(nums []int, target int) {
	j := 0 // 用来记录非 target 元素的记录
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			nums[j] = nums[i]

			if i != j {
				nums[i] = target
			}

			j++
		}

	}
}

func main() {
	x := []int{5, 1, 6, 1}
	MoveTarget(x, 1)
	fmt.Println(x)
}
