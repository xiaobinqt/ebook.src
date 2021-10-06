package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("i am dead")
	println("exit...")
}
