package main

import (
	"fmt"
	"log"
	"os"
)

//静态链表节点
type NodeStr struct {
	data   string
	cursor int
}

const maxSize int = 10

//初始化链表
func initList(size int) []NodeStr {
	if size < 3 {
		log.Fatal("size参数错误")
		return nil
	}
	list := make([]NodeStr, size)
	for i := 0; i < size-2; i++ {
		list[i].cursor = i + 1
	}
	list[size-2].cursor = 0
	list[size-1].cursor = 0
	return list
}

//显示链表结构
func traverse(list []NodeStr) {
	for _, v := range list {
		fmt.Printf("%5d", v.cursor)
	}
	fmt.Println()
	for _, v := range list {
		fmt.Printf("%5s", v.data)
	}
	fmt.Println()
	for i, _ := range list {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
}

//回收链表到备用链表
func destroyList(list []NodeStr) {
	if list[maxSize-1].cursor == 0 {
		return
	}
	j := list[maxSize-1].cursor
	list[maxSize-1].cursor = 0
	i := list[0].cursor
	list[0].cursor = j
	if j > 0 {
		j = list[j].cursor
	}
	list[j].cursor = i
}

//判断是否为空
func isempty(list []NodeStr) bool {
	if list[maxSize-1].cursor == 0 {
		return true
	}
	return false
}

//链表长度
func length(list []NodeStr) int {
	i, j := 0, list[maxSize-1].cursor
	for j > 0 {
		j = list[j].cursor
		i++
	}
	return i
}

//获取指定位置的节点数据
func getElement(list []NodeStr, index int) string {
	if index < 1 || index > maxSize-2 {
		log.Fatal("index out of range")
		return ""
	}
	i := list[maxSize-1].cursor
	j := 1
	for i > 0 && j < index {
		j++
		i = list[i].cursor
	}
	if j != index {
		return ""
	}
	return list[i].data
}

//获取数据元素的位置
func locateElem(list []NodeStr, data string) int {
	locate := 0
	i := list[maxSize-1].cursor
	for i > 0 {
		locate++
		if list[i].data == data {
			return locate
		}
		i = list[i].cursor
	}
	return locate
}

//获取元素的前驱节点
func priorElem(list []NodeStr, data string) string {
	if isempty(list) {
		return ""
	}
	i := list[maxSize-1].cursor
	var j int
	for i > 0 {
		j = list[i].cursor
		if list[j].data == data {
			return list[i].data
		}
		i = j
	}
	return ""
}

//获取元素的后驱节点
func nextElem(list []NodeStr, data string) string {
	if isempty(list) {
		return ""
	}
	i := list[maxSize-1].cursor
	var j int
	for i > 0 {
		j = list[i].cursor
		if list[i].data == data {
			return list[j].data
		}
		i = j
	}
	return ""
}

//分配节点
func malloc(list []NodeStr) int {
	i := list[0].cursor
	if i == 0 {
		os.Exit(0)
	}
	list[0].cursor = list[i].cursor
	return i
}

//回收节点
func free(list []NodeStr, index int) {
	list[index].cursor = list[0].cursor
	list[0].cursor = index
}

//插入节点
func insertLIst(list []NodeStr, index int, data string) {
	if index < 1 || index > length(list) {
		os.Exit(0)
	}
	i := list[maxSize-1].cursor
	j := 1
	for i > 0 && j < index-1 {
		j++
		i = list[i].cursor
	}
	tmp := list[i].cursor
	cur := malloc(list)
	list[cur].data = data
	list[cur].cursor = tmp
	list[i].cursor = cur
}

//删除节点
func deleteList(list []NodeStr, index int) string {
	if index < 1 || index > length(list) {
		return "删除参数错误"
	}
	i := list[maxSize-1].cursor
	j := 1
	for i > 0 && j < index-1 {
		j++
		i = list[i].cursor
	}
	tmp := list[i].cursor
	list[i].cursor = list[tmp].cursor
	val := list[tmp].data
	free(list, tmp)
	return val
}

//遍历链表
func traveList(list []NodeStr) {
	if list[maxSize-1].cursor == 0 {
		return
	}
	i := list[maxSize-1].cursor
	j := 1
	for i > 0 {
		fmt.Printf("第%d个节点为：%s\n", j, list[i].data)
		i = list[i].cursor
		j++
	}
}

func main() {
	var list []NodeStr
	list = initList(maxSize)
	list[1].data = "A"
	list[9].cursor = 1
	list[2].data = "C"
	list[2].cursor = 0
	list[0].cursor = 3
	//traverse(list)
	insertLIst(list, 2, "B")
	traveList(list)
	reval := deleteList(list, 1)
	fmt.Println("删除的节点为：", reval)
	traveList(list)
}
