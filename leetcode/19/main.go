package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	fast := dummy
	low := dummy
	for fast != nil && n != 0 {
		fast = fast.Next
		n--
	}
	fast = fast.Next

	for fast != nil {
		fast = fast.Next
		low = low.Next
	}

	low.Next = low.Next.Next
	return dummy.Next
}

func main() {
	ret := removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}, 2)
	fmt.Println(ret)
}
