package main

func test6() (z int) {
	defer func() {
		println("defer :", z)
		z += 100
	}()

	return 100
}

func main() {
	println("test6 : ", test6())
}
