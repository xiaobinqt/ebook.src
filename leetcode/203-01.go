package main

import (
	"fmt"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := dummyHead
	fmt.Println(dummyHead, curr)
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummyHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	x := removeElements(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}, 1)

	fmt.Println(x)
}
