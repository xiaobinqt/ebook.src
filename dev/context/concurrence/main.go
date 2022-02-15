package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "wb", "test01")
	go func() {
		for {
			_ = context.WithValue(ctx, "wb", "test02")
		}
	}()
	go func() {
		for {
			_ = context.WithValue(ctx, "wb", "test03")
		}
	}()
	go func() {
		for {
			fmt.Println(ctx.Value("wb"))
		}
	}()
	go func() {
		for {
			fmt.Println(ctx.Value("wb"))
		}
	}()
	time.Sleep(10 * time.Second)
}
