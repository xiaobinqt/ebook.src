package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	jsFilename := "node1.js"
	xx := ""
	if len(strings.Split(jsFilename, ".")) > 1 {
		xx = strings.Split(jsFilename, ".")[0]
	}
	fmt.Println(xx)
}
