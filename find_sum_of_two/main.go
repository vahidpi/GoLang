package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	list := []int{5, 7, 1, 2, 8, 4, 3}
	A := 10
	x := findSumOfTwo(A, list)
	fmt.Printf("Some of two numbers from %v is equal %d : %v \n", list, A, x)
}

func findSumOfTwo(a int, list []int) bool {
	var sub = make(map[int]bool)

	for _, s := range list {
		z := a - s
		if sub[z] == true {
			return true
		}
		sub[s] = true
	}
	return false
}
