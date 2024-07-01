package main

import "fmt"

/**
3 个函数分别打印
cat、dog、fish，
要求每个函数都要起一个 goroutine，

按照 cat、dog、fish 顺序打印在屏幕上 100 次。


输入一个链表，输出该链表中倒数第k个节点。
为了符合大多数人的习惯，本题从1开始计数，
即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，
从头节点开始，它们的值依次是1、2、3、4、5、6。
这个链表的倒数第3个节点是值为4的节点。

示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.


*/

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func k(head *LinkNode, n int) *LinkNode {
	if head == nil {
		return head
	}

	leng := 0
	curr := head
	for curr != nil {
		leng++
		curr = curr.Next
	}

	cut := leng - n
	c2 := head
	for i := 0; i < cut; i++ {
		c2 = c2.Next
	}

	return c2
}

func main() {
	head := &LinkNode{
		Val: 1,
		Next: &LinkNode{
			Val: 2,
			Next: &LinkNode{
				Val: 3,
				Next: &LinkNode{
					Val: 4,
					Next: &LinkNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	x := k(head, 2)
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}
