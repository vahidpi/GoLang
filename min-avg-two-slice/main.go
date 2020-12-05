package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	A := []int{4, 2, 2, 1, 1, 1, 8}
	indexNumber := solution(A)
	fmt.Printf("The start index of Min Avg Two Slice is %d \n", indexNumber)

}

func solution(A []int) int {

	var avg float64 = 100000
	index := -1

	for i := 0; i < len(A)-1; i++ {
		tmpAvg := float64(A[i]+A[i+1]) / 2
		if tmpAvg < avg {
			index = i
			avg = tmpAvg
		}
	}

	return index
}

func solutionMinAvg(A []int) int {
	var avg float64 = 100000
	index := -1
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			tmpAvg := float64(sum(A[i:j])) / float64(j-i+1)
			if tmpAvg < avg {
				avg = tmpAvg
				index = i
			}
		}
	}
	return index
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
