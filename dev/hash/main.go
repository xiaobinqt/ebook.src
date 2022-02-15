package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type Stu struct {
}

func main() {
	var s *Stu
	fmt.Println(s)
}
