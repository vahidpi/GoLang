package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Start")
	max := solution("34539849")
	fmt.Printf("Max 2 digits number is %d \n", max)
}

func solution(S string) int {
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
