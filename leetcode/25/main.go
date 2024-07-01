package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 反转链表
	reverse := func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		var (
			prev, curr, next *ListNode
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

	dummy := &ListNode{}
	var (
		prev, end *ListNode
	)

	dummy.Next = head
	prev = dummy
	end = dummy

	for end.Next != nil {
		// 把 end 往后移动 k 个位置
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start := prev.Next
		next := end.Next
		end.Next = nil

		prev.Next = reverse(start)
		start.Next = next

		prev = start
		end = start

	}

	return dummy.Next
}

func main() {
	ret := reverseKGroup(&ListNode{
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
