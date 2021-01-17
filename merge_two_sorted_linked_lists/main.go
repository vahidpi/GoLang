package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Start")

	var sortedList1 list.List
	var sortedList2 list.List
	// list1 := []int{4, 9, 15, 16, 24}
	list1 := []int{3, 7, 12, 16}
	for _, num := range list1 {
		sortedList1.PushBack(num)
	}
	// list2 := []int{7, 8, 10, 19, 20}
	list2 := []int{6, 7, 8, 9}
	for _, num := range list2 {
		sortedList2.PushBack(num)
	}

	mergedList := mergeSortedList(sortedList1, sortedList2)
	fmt.Print("Merged list is : ")
	for e := mergedList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}
	fmt.Println()
}

func mergeSortedList(list1 list.List, list2 list.List) list.List {
	var outList list.List
	head1 := list1.Front()
	head2 := list2.Front()

	for (head1 != nil) && (head2 != nil) {
		if head1.Value.(int) < head2.Value.(int) {
			outList.PushBack(head1.Value)
			head1 = head1.Next()
		} else {
			outList.PushBack(head2.Value)
			head2 = head2.Next()
		}
	}

	for head1 != nil {
		outList.PushBack(head1.Value)
		head1 = head1.Next()
	}

	for head2 != nil {
		outList.PushBack(head2.Value)
		head2 = head2.Next()
	}

	return outList
}
