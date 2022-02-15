package main

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

var studentPool = &sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

func New(name string, age int) *student {
	stu := studentPool.Get().(*student)
	stu.Name = name
	stu.Age = age
	return stu
}

func Release(stu *student) {
	stu.Name = ""
	stu.Age = 0
	studentPool.Put(stu)
}

func test() {
	stu := New("tom", 30)
	defer Release(stu)

	fmt.Println(stu)
}

func main() {
	test()
}
