package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	s := "aabbffddeaee"
	minNumber := solution(s)
	fmt.Printf("Min Deletions to Make Frequency of Each Letter Unique for '%s' is %d \n", s, minNumber)
}

func solution(s string) int {

	listOfChar := make(map[rune]int)
	for _, char := range s {
		listOfChar[char] = listOfChar[char] + 1
	}

	count := 0
	var listTarget []int
	for _, val := range listOfChar {
		a := val
		for a != 0 && isInList(a, listTarget) {
			count++
			a = a - 1
		}
		if a != 0 {
			listTarget = append(listTarget, a)
		}
	}

	return count
}

func isInList(value int, list []int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
