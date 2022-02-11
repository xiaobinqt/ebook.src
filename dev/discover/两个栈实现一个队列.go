package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

//构造函数
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

//尾部插入
func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

//头部删除,back函数返回其list最尾部的值
func (this *CQueue) DeleteHead() int {
	//如果第二个栈为空
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func main() {
	var cq CQueue
	cq = Constructor()
	cq.AppendTail(1)
	fmt.Println(cq.DeleteHead())
}
