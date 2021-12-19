package main

import (
	"fmt"
	"os"
)

func pac() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	f()

	return
}

func f2() {
	fmt.Println("222222222222222")
}

func main() {
	//var done = make(chan struct{})
	throwsPanic(pac)

	f2()
	//<-done
}
