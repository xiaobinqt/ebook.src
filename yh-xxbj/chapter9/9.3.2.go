package main

var x = 100

func init() {
	println("init: ", x)
	x++
}

func main() {
	println("main: ", x)
}
