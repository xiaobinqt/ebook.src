package main

import "fmt"

type LinkNode struct {
	Next *ListNode
	Val  int
}

/**
1 -> 2 -> 3
*/
func fz(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	var (
		prev, next *ListNode
	)

	curr := dummy.Next
	for curr.Next != nil {
		next = curr.Next
		curr.Next = prev
		curr.Next = curr

		prev = curr.Next
		curr = next
	}

	return dummy.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	curr := dummy
	for curr.Next != nil {
		next := curr.Next
		if next.Val == val {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}

func main() {
	ret := fz(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	fmt.Println(ret)
}
