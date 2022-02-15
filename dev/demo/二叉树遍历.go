package main

import "fmt"

type TreeNode struct {
	Value byte
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(node *TreeNode, values []string) []string {
	if node != nil {
		values = append(values, string(node.Value))
		values = preorderTraversal(node.Left, values)
		values = preorderTraversal(node.Right, values)
	}

	return values
}

// 中序遍历
func midOrderTraversal(node *TreeNode, values []string) []string {
	if node != nil {
		values = preorderTraversal(node.Left, values)
		values = append(values, string(node.Value))
		values = preorderTraversal(node.Right, values)
	}

	return values
}

// 后序遍历
func postOrderTraversal(node *TreeNode, values []string) []string {
	if node != nil {
		values = preorderTraversal(node.Left, values)
		values = preorderTraversal(node.Right, values)
		values = append(values, string(node.Value))
	}

	return values
}

// 按层遍历
func LevelOrderTraversal(node *TreeNode, values []string) []string {
	if node != nil {
		// 采用队列实现
		queue := make([]*TreeNode, 0)
		queue = append(queue, node) // queue push
		for len(queue) > 0 {
			node = queue[0]
			//fmt.Printf("%d -> ", tree.Value)
			values = append(values, string(node.Value))

			queue = queue[1:] // queue pop
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return values
}

func main() {
	node := &TreeNode{
		Value: 'A',
		Left: &TreeNode{
			Value: 'B',
			Left: &TreeNode{
				Value: 'D',
				Left: &TreeNode{
					Value: 'H',
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Value: 'I',
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Value: 'E',
				Left: &TreeNode{
					Value: 'J',
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Value: 'C',
			Left: &TreeNode{
				Value: 'F',
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Value: 'G',
				Left:  nil,
				Right: nil,
			},
		},
	}

	values := make([]string, 0)
	fmt.Println("前序: ", preorderTraversal(node, values))
	fmt.Println("中序: ", midOrderTraversal(node, values))
	fmt.Println("后序: ", postOrderTraversal(node, values))
	fmt.Println("按层: ", LevelOrderTraversal(node, values))
}
