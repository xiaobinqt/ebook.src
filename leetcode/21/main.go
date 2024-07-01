package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (x *ListNode) {
	dummy := &ListNode{}
	prev := dummy

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			prev.Next = list2
			list2 = list2.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
		}
		prev = prev.Next
	}

	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}

	return dummy.Next
}

func main() {
	x := mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}}, &ListNode{Val: 1, Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}})
	fmt.Println(x)
}
