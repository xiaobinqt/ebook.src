package main

import (
	"fmt"
	"io"
)

type MyReader struct {
	Name string
}

func (r MyReader) Read(p []byte) (n int, err error) {

	return
}

func main() {
	var reader io.Reader
	reader = &MyReader{
		Name: "a.txt",
	}
	fmt.Println(reader)
}
