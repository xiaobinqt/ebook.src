package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		tmp := curr.Next
		tmp1 := curr.Next.Next.Next

		curr.Next = curr.Next.Next
		curr.Next.Next = tmp
		curr.Next.Next.Next = tmp1

		curr = curr.Next.Next
	}

	return dummy.Next
}

func main() {
	x := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	})
	fmt.Println(x)
}
