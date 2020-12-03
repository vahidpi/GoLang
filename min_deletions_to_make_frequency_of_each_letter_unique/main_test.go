package main

import "testing"

func TestSolution(t *testing.T) {

	var tests = []struct {
		value string
		want  int
	}{
		{"eeeeffff", 1},
		{"aabbffddeaee", 6},
		{"llll", 0},
		{"example", 4},
	}

	for _, tt := range tests {
		gotValue := solution(tt.value)

		if gotValue != tt.want {
			t.Errorf("got %d, want %d", gotValue, tt.want)
		}
	}
}
