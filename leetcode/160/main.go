package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var (
		currA, currB = headA, headB
		lenA, lenB   int
	)
	// 计算链表的长度
	for currA != nil {
		lenA++
		currA = currA.Next
	}

	for currB != nil {
		lenB++
		currB = currB.Next
	}

	currA, currB = headA, headB

	// 判断谁是最长的,把最长的赋值给 currA
	gap := 0
	if lenA >= lenB {
		gap = lenA - lenB
	} else {
		gap = lenB - lenA
		currA, currB = currB, currA
	}

	for gap > 0 { // 移动最大的位置
		currA = currA.Next
		gap--
	}

	// 同时移动
	for currA != nil {
		if currA == currB {
			return currA // 这里 return currB 也行
		}
		currA = currA.Next
		currB = currB.Next
	}

	return nil
}

func main() {
	t := &ListNode{
		Next: &ListNode{
			Next: nil,
			Val:  9,
		},
		Val: 8,
	}

	ca := &ListNode{
		Next: &ListNode{
			Next: t,
			Val:  2,
		},
		Val: 1,
	}

	cb := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: t,
				Val:  5,
			},
			Val: 4,
		},
		Val: 3,
	}
	ret := getIntersectionNode(ca, cb)
	fmt.Println(ret)
}
