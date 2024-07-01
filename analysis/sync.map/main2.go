package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	// 1. 写入
	m.Store("qcrao", nil)
	m.Store("stefno", nil)

	// 2. 读取
	r1, err := m.Load("qcrao")
	fmt.Println(r1, err)

}
