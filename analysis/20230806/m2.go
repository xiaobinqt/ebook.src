package main

import "fmt"

func f3() {
	var err error
	defer func() {
		fmt.Println(err)
	}()

	err = fmt.Errorf("111")
	return
}

func main() {
	f3()
}
