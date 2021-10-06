package main

import (
	"fmt"
	"log"
)

type DivError struct {
	x, y int
}

func (d DivError) Error() string {
	return "division by zero"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{
			x: x,
			y: y,
		}
	}

	return x / y, nil
}

func main() {
	z, err := div(5, 0)

	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Println("111", e, e.x, e.y)
		default:
			fmt.Println("default: ", e)
		}

		log.Fatalln("hhhh", err)
	}

	println(z)
}
