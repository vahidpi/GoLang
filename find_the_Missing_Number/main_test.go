package main

import "testing"

func TestFindX(t *testing.T) {

	var tests = []struct {
		value []int
		want  int
	}{
		{[]int{1, 3, 5, 4, 6}, 2},
		{[]int{2, 3, 5, 4, 6}, 1},
		{[]int{4, 3, 7, 1, 2, 6}, 5},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 6},
	}

	for _, tt := range tests {
		gotValue := findX(tt.value)

		if gotValue != tt.want {
			t.Errorf("got %d, want %d", gotValue, tt.want)
		}
	}
}
