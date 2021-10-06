package main

func main() {
	x := 10
	p := &x

	*p++
	println(p, x)

}
