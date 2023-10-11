package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[int32]int)
	for _, w := range s {
		if _, ok := count[w]; ok {
			count[w]++
		} else {
			count[w] = 1
		}
	}

	fmt.Println(count)

	for _, w := range t {
		if _, ok := count[w]; ok {
			count[w]--
		}
	}

	fmt.Println(count)

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(string(int32(97)))
}
