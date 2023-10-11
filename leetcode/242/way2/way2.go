package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [30]int

	for i, _ := range s {
		arr[s[i]-'a']++
	}

	for i, _ := range t {
		arr[t[i]-'a']--
	}

	fmt.Println(arr)
	for _, e := range arr {
		if e != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("rat", "car"))
}
