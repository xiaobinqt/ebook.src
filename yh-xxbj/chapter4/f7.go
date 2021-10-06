package main

func test5() []func() {
	var s = make([]func(), 0)

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			//println(&i, i)
			println(&x, x)
		})
	}

	return s
}

func main() {
	for _, f := range test5() {
		f()
	}
}
