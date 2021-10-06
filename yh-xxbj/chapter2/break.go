package main

func main() {
outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				println()
				continue outer
			}

			if x > 2 {
				break outer
			}

			println(x, ":", y, " ")
		}
	}
}
