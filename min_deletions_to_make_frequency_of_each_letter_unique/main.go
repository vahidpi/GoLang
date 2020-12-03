package main

import (
	"fmt"
	"sort"
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

	if len(listOfChar) <= 1 {
		return 0
	}
	if len(listOfChar) == 2 {
		return 1
	}

	var list []int
	for _, val := range listOfChar {
		list = append(list, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	count := 0
	var listTarget []int
	for i := 0; i < len(list); i++ {
		a := list[i]
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
