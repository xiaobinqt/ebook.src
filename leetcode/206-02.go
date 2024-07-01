package main

import (
	"fmt"

	"go.src/leetcode/model"
)

func reverseList2(head *model.ListNode) *model.ListNode {
	var (
		next, prev, curr *model.ListNode
	)

	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

func main() {
	ret := reverseList2(&model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val: 2,
			Next: &model.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	fmt.Println(ret)
}
