package main

import (
	"fmt"
	"log"
)

type TestError struct {
}

func (*TestError) Error() string {
	return "error.."
}

func test(x int) (int, error) {
	var err *TestError

	if x < 0 {
		err = new(TestError)
		x = 0
	} else {
		x += 100
	}

	fmt.Printf("%#v, %+v \n", err, err)
	return x, err
	//return x, nil
}

func main() {
	x, err := test(100)
	if err != nil {
		log.Fatalln("err != nil")
	}

	println(x)
}
