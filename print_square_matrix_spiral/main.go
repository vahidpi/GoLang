package main

import (
	"fmt"
)

func main() {

	data3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	printMat(data3)
}

func printMat(mat [][]int) {

	n := len(mat)

	if n == 1 {
		fmt.Print(mat[0][0])
		return
	}

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Print(mat[0][i])
	}
	for i := 1; i < n; i++ {
		fmt.Print(mat[i][n-1])
	}
	for i := n - 2; i >= 0; i-- {
		fmt.Print(mat[n-1][i])
	}
	for i := n - 2; i >= 1; i-- {
		fmt.Print(mat[i][0])
	}

	if n > 2 {
		var tmp = make([][]int, len(mat)-2)
		for i := range tmp {
			tmp[i] = make([]int, len(mat)-2)
		}

		for i := 1; i < n-1; i++ {
			for j := 1; j < n-1; j++ {
				tmp[i-1][j-1] = mat[i][j]
			}
		}
		printMat(tmp)
	}

	return
}
