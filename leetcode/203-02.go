package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
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
