package main

import "fmt"

type binarySearchTree struct {
	value       int
	left, right *binarySearchTree
}

//初始化并添加根节点的值
func NewBinarySearchTree(rootValue int) *binarySearchTree {
	return &binarySearchTree{value: rootValue}
}

//添加元素
func (t *binarySearchTree) Insert(value int) *binarySearchTree {
	if t == nil {
		t = NewBinarySearchTree(value)
		return t
	}
	if value < t.value {
		t.left = t.left.Insert(value)
	} else {
		t.right = t.right.Insert(value)
	}
	return t
}

//是否包含某一个元素
func (t *binarySearchTree) Contains(value int) bool {
	if t == nil {
		return false
	}
	v := t.compareTo(value)

	if v < 0 {
		return t.left.Contains(value)
	} else if v > 0 {
		return t.right.Contains(value)
	} else {
		return true
	}
}

func (t *binarySearchTree) compareTo(value int) int {
	return value - t.value
}

//移除元素
func (t *binarySearchTree) Remove(value int) *binarySearchTree {
	if t == nil {
		return t
	}
	compareResult := t.compareTo(value)
	if compareResult < 0 {
		t.left = t.left.Remove(value)
	} else if compareResult > 0 {
		t.right = t.right.Remove(value)
	} else if t.left != nil && t.right != nil {
		t.value = t.right.FindMin()
		t.right = t.right.Remove(t.value)
	} else if t.left != nil {
		t = t.left
	} else {
		t = t.right
	}
	return t
}

//查找最大值
func (t *binarySearchTree) FindMax() int {
	if t == nil {
		fmt.Println("tree is empty")
		return -1
	}
	if t.right == nil {
		return t.value
	} else {
		return t.right.FindMax()
	}
	//也可以用下面的方法
	//return t.FindMaxNode().value
}

//查找最大的节点
func (t *binarySearchTree) FindMaxNode() *binarySearchTree {
	if t != nil {
		for t.right != nil {
			t = t.right
		}
	}
	return t
}

//查找最小值
func (t *binarySearchTree) FindMin() int {
	if t == nil {
		fmt.Println("tree is empty")
		return -1
	}
	if t.left == nil {
		return t.value
	} else {
		return t.left.FindMin()
	}
	//也可以直接用下面的方法
	//return t.FindMinNode().value
}

//查找最小的节点
func (t *binarySearchTree) FindMinNode() *binarySearchTree {
	if t != nil {
		for t.left != nil {
			t = t.left
		}
	}

	return t
}

//获取树种所有的元素值（并按从小到大排序）
func (t *binarySearchTree) GetAll() []int {
	values := []int{}
	return appendValues(values, t)
}

func appendValues(values []int, t *binarySearchTree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	binaryTree := NewBinarySearchTree(50)
	binaryTree.Insert(20)
	binaryTree.Insert(10)
	binaryTree.Insert(100)
	binaryTree.Insert(60)
	binaryTree.Insert(70)
	binaryTree.Insert(5)
	binaryTree.Insert(35)
	binaryTree.Insert(40)
	fmt.Println(binaryTree.GetAll())

	fmt.Println(binaryTree.Contains(30))
	fmt.Println(binaryTree.Contains(20))

	fmt.Println(binaryTree.FindMin())
	fmt.Println(binaryTree.FindMinNode().value)

	fmt.Println(binaryTree.FindMax())
	fmt.Println(binaryTree.FindMaxNode().value)

	binaryTree.Remove(20)
	fmt.Println(binaryTree.GetAll())

}
