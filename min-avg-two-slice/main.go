package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	A := []int{4, 2, 2, 5, 1, 5, 8}
	indexNumber := solution(A)
	fmt.Printf("The start index of Min Avg Two Slice is %d \n", indexNumber)
}

func solution(A []int) int {

	if len(A) == 2 {
		return 0
	}

	minSliceTwo := A[0] + A[1]
	minTwoIndex := 0

	minSliceThree := 100000
	minThreeIndex := 0

	for i := 2; i < len(A); i++ {
		sliceTwo := A[i-1] + A[i]
		if sliceTwo < minSliceTwo {
			minSliceTwo = sliceTwo
			minTwoIndex = i - 1
		}

		sliceThree := sliceTwo + A[i-2]
		if sliceThree < minSliceThree {
			minSliceThree = sliceThree
			minThreeIndex = i - 2
		}
	}
	averageMinTwo := minSliceTwo * 3
	averageMinThree := minSliceThree * 2

	if averageMinTwo == averageMinThree {
		if minTwoIndex < minThreeIndex {
			return minTwoIndex
		}
		return minThreeIndex

	}
	if averageMinTwo < averageMinThree {
		return minTwoIndex
	}
	return minThreeIndex

	//return 0
}
