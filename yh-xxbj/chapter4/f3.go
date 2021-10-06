package main

func test1(p *int) {
	go func() {
		println(p)
	}()
}

func main() {
	x := 100
	p := &x
	test1(p)
}
