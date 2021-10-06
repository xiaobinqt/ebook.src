package main

import "log"

func test7() {
	defer println("test.1")
	defer println("test.2")

	panic("i am dead")
}

func main() {
	defer func() {
		log.Println(recover())
	}()

	test7()
}
