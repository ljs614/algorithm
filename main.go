package main

import (
	"fmt"

	"github.com/ljs614/algorithm/sort"
)

func main() {
	list := []int{1, 10, 5, 8, 7, 6, 4, 3, 2, 9}
	// sort.SelectionSort(list)
	// sort.BubbleSort(list)
	// sort.InsertionSort(list)
	// sort.QuickSort(list, 0, len(list)-1)
	// sort.QuickSortDesc(list, 0, len(list)-1)
	sort.MergeSort(list, 0, len(list)-1)
	fmt.Println(list)
}
