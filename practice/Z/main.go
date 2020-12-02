package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("ok")

	A := [6]int{5, 2, 4, 6, 3, 7}

	a := Solutionx(A)
	fmt.Println(a)

}

func Solutio(A [6]int) int {
	result := 0
	N := len(A)
	var prefix [7]int
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + A[i-1]
	}
	var avg float32 = 100000
	for i := 1; i < N; i++ {
		for j := i + 1; j <= N; j++ {
			var temp float32
			temp = float32(prefix[j]-prefix[i-1]) / float32((j - i + 1))
			if temp < avg {
				avg = temp
				result = i - 1
			}
		}
	}
	return result
}
func Solution(A [6]int) int {
	minAvgIdx := 0
	var minAvgVal float64
	minAvgVal = float64(A[0]+A[1]) / 2
	var currAvg float64
	for i := 0; i < len(A); i++ {

		currAvg = float64(A[i] + A[i+1]/2)
		if currAvg < minAvgVal {
			minAvgVal = currAvg
			minAvgIdx = i
		}

		currAvg = float64((A[i] + A[i+1] + A[i+2])) / 3
		if currAvg < minAvgVal {
			minAvgVal = currAvg
			minAvgIdx = i
		}
	}

	currAvg = ((float64)(A[len(A)-2] + A[len(A)-1])) / 2
	if currAvg < minAvgVal {
		minAvgVal = currAvg
		minAvgIdx = len(A) - 2
	}
	return minAvgIdx
}

func Solutionxx(A [6]int) int {
	if len(A) < 2 {
		return -1
	}

	result := 0
	minAvg := float64(A[0]+A[1]) / 2
	var curAvg float64
	for index := 0; index < len(A)-2; index++ {
		curAvg = float64(A[index]+A[index+1]) / 2
		if curAvg < minAvg {
			minAvg = curAvg
			result = index
		}

		curAvg = float64(A[index]+A[index+1]+A[index+2]) / 3
		if curAvg < minAvg {
			minAvg = curAvg
			result = index
		}
	}

	curAvg = float64(A[len(A)-2]+A[len(A)-1]) / 2
	if curAvg < minAvg {
		minAvg = curAvg
		result = len(A) - 2
	}

	return result
}

func Solutionx(A [6]int) int {
	// write your code in Go 1.4

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

}

func Solution2(S string) int {

	var bank [26]int
	strLen := len(S)
	for i := 0; i < strLen; i++ {
		x := S[i]
		bank[x-'a']++
	}
	min := 300000

	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {

			z := bank[i] + bank[j]
			if strLen-z < min {
				min = strLen - z
			}
		}
	}
	if min == 0 && countNonZero(bank) != 0 {
		return 1
	}
	return min
}

func countNonZero(list [26]int) int {
	c := 0
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			c++
		}
	}
	return c
}

func Solution1(S string) int {
	// write your code in Go 1.4
	max := 0
	for i := 0; i < len(S)-1; i++ {
		num, _ := strconv.Atoi(S[i : i+2])
		if num > max {
			max = num
		}
	}
	return max
}
