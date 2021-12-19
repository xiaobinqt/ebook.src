package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result = make([]int, 0)

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd)
	fmt.Println("odd elements of slice are: ", odd)

	even := filter(slice, isEven)
	fmt.Println("even elements of slice are: ", even)
}
