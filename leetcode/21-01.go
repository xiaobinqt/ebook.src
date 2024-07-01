package main

import (
	"fmt"

	"go.src/leetcode/model"
)

/**
参考：
递归实现
https://leetcode.cn/problems/merge-two-sorted-lists/solution/yi-kan-jiu-hui-yi-xie-jiu-fei-xiang-jie-di-gui-by-/


*/

func mergeTwoLists(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

/**
[1,2,4], list2 = [1,3,4]
*/

func PrintLinkedList(l *model.ListNode) {
	var res = make([]int, 0)
	for {
		res = append(res, l.Val)
		if l.Next == nil {
			break
		}
		l = l.Next
	}

	fmt.Println(res)
}

func main() {
	list1 := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val: 2,
			Next: &model.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &model.ListNode{
		Val: 1,
		Next: &model.ListNode{
			Val: 3,
			Next: &model.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	PrintLinkedList(mergeTwoLists(list1, list2))
}
