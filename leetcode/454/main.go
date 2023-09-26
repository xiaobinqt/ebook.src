package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			target := 0 - (c + d)
			if v, ok := m[target]; ok {
				count += v
			}
		}
	}

	return count
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	ret := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Println(ret)
}
