package main

import "runtime/debug"

func test8() {
	panic("i am dead")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	test8()
}
