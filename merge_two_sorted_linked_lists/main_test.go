package main

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedList(t *testing.T) {

	var tests = []struct {
		fValue []int
		sValue []int
		want   []int
	}{
		{[]int{1, 3, 5, 6}, []int{4, 9, 15, 16, 24}, []int{1, 3, 4, 5, 6, 9, 15, 16, 24}},
		{[]int{2, 3, 4, 6}, []int{3, 4}, []int{2, 3, 3, 4, 4, 6}},
		{[]int{3, 7, 12, 16}, []int{6, 7, 8, 9}, []int{3, 6, 7, 7, 8, 9, 12, 16}},
		{[]int{3, 7, 10, 12, 18, 24, 25}, []int{1, 2, 15, 16}, []int{1, 2, 3, 7, 10, 12, 15, 16, 18, 24, 25}},
		{[]int{4, 9, 15, 16, 24}, []int{7, 8, 10, 19, 20}, []int{4, 7, 8, 9, 10, 15, 16, 19, 20, 24}},
	}

	for _, tt := range tests {
		var sortedList1 list.List
		var sortedList2 list.List
		var sortedList3 []int
		var gotValue list.List

		list1 := tt.fValue
		for _, num := range list1 {
			sortedList1.PushBack(num)
		}
		list2 := tt.sValue
		for _, num := range list2 {
			sortedList2.PushBack(num)
		}

		gotValue = mergeSortedList(sortedList1, sortedList2)
		head := gotValue.Front()

		for head != nil {
			a := head.Value.(int)
			sortedList3 = append(sortedList3, a)
			head = head.Next()
		}

		if !assert.Equal(t, tt.want, sortedList3) {
			t.Errorf("Expect %v, got %v", tt.want, sortedList3)
		}
	}
}
