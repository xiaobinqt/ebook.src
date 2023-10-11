package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	low, fast := head, head

	// 因为快指针走 2 步，不要判断 next 是否为 nil
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if fast == low {
			for low != head {
				low = low.Next
				head = head.Next

			}
			return head
		}
	}

	return nil
}

func main() {
	index2 := &ListNode{
		Val:  101,
		Next: nil,
	}
	index := &ListNode{
		Val:  100,
		Next: index2,
	}
	index.Next.Next = &ListNode{
		Val:  4,
		Next: index,
	}
	fmt.Println(detectCycle(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: index,
			},
		},
	}))
}
