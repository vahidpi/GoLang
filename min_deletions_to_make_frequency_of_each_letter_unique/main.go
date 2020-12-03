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
		if !isInList(a, listTarget) {
			listTarget = append(listTarget, a)
		} else {
			de := true
			for j := a - 1; de; j-- {
				if j >= 0 {
					count++
					if !isInList(j, listTarget) {
						listTarget = append(listTarget, j)
						de = false
					}
				} else {
					de = false
				}
			}
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
