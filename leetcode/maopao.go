package main

import (
	"fmt"
	"strings"
)

func maopao(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 从小到大排序
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp

			}
		}
	}

	return nums
}

func main() {
	fmt.Println(maopao([]int{1, 10, 8, 100}))
}

// update_time,desc|age,asc
func FormatOrderBy(order string, aliasTable ...string) string {
	if len(aliasTable) > 0 { // 只取第一个
		alias := aliasTable[0]
		orderArr := strings.Split(order, "|")
		for idx, each := range orderArr {
			orderArr[idx] = fmt.Sprintf("%s.%s", alias, each)
		}
		order = strings.Join(orderArr, "|")

		order = strings.ReplaceAll(order, ",", " ")
		return strings.ReplaceAll(order, "|", ",")

	}
	order = strings.ReplaceAll(order, ",", " ")
	return strings.ReplaceAll(order, "|", ",")
}
