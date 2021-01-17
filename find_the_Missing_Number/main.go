package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	list := []int{3, 7, 1, 2, 8, 4, 5}
	x := findX(list)
	fmt.Printf("Missing number is %d \n", x)
}

func findX(list []int) int {
	n := len(list) + 1
	sumKey := 0
	sumValue := 0
	for i, s := range list {
		sumKey += i + 1
		sumValue += s

	}
	sumKey += n
	return sumKey - sumValue
}
