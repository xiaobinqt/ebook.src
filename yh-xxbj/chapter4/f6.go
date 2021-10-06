package main

func test4(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test4(123)
	f()
}
