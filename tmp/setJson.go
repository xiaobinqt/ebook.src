package main

import (
	"fmt"
	"time"

	"go.src/tmp/qworker"
)

func main() {
	qworker.NewPhysicalDepApp()

	f := qworker.Handler{
		F: func(args ...interface{}) error {
			time.Sleep(5 * time.Second)
			fmt.Println("111111", args[0], args[1])
			return nil
		},
		Args: []interface{}{1, 2},
	}
	qworker.DepApp.AddAction(f)

	f2 := qworker.Handler{
		F: func(args ...interface{}) error {
			fmt.Println("222222", args[0], args[1])
			return nil
		},
		Args: []interface{}{3, 4},
	}
	qworker.DepApp.AddAction(f2)

	for {

	}
}
