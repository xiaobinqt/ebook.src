package main

import "fmt"

func main() {
	reverseString := func(s string) string {
		r := []rune(s)
		for from, to := 0, len(r)-1; from < to; from, to = from+1, to-1 {
			r[from], r[to] = r[to], r[from]
		}
		return string(r)
	}

	fmt.Println(reverseString("abc"))
}
