package main

import (
	"fmt"

	"github.com/ljs614/algorithm/sort"
)

func main() {
	list := []int{1, 10, 5, 8, 7, 6, 4, 3, 2, 9}
	sort.SelectionSort(list)
	fmt.Println(list)
}
