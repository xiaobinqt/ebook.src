package main

import "fmt"

type NodeInt struct {
	Data int
	Next *NodeInt
}

func New() *NodeInt {
	return &NodeInt{
		Data: 0,
		Next: nil,
	}
}

func (l *NodeInt) GetElem(i int) (data int, err error) {
	p := l.Next
	for j := 1; j < i; j++ {
		if p == nil {
			// 返回错误
			return -100001, fmt.Errorf("not found")
		}
		p = p.Next
	}

	return p.Data, nil
}

func (l *NodeInt) Insert(i int, data int) bool {
	p := l
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}

	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &NodeInt{Data: data}
	s.Next = p.Next
	p.Next = s
	return true
}

func main() {
	p := New()
	p.Next = &NodeInt{
		Data: 1111,
		Next: &NodeInt{
			Data: 2222,
			Next: &NodeInt{
				Data: 3333,
				Next: nil,
			},
		},
	}

	x, err := p.GetElem(2)
	if err != nil {
		fmt.Println("getElem err ", err.Error())
		return
	}

	fmt.Println("getElem success: ", x)

	insertSucc := p.Insert(3, 11)
	fmt.Println("insertSucc : ", insertSucc)
}
