package main

import "testing"

func TestFindSumOfTwo(t *testing.T) {

	var tests = []struct {
		list  []int
		value int
		want  bool
	}{
		{[]int{5, 7, 1, 2, 8, 4, 3}, 10, true},
		{[]int{2, 3, 5, 4, 6}, 7, true},
		{[]int{4, 3, 7, 6}, 5, false},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 20, false},
	}

	for _, tt := range tests {
		gotValue := findSumOfTwo(tt.value, tt.list)

		if gotValue != tt.want {
			t.Errorf("for %v and %v ,got %v, want %v", tt.value, tt.list, gotValue, tt.want)
		}
	}
}
