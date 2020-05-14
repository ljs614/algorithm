package sort

//SelectionSort ...
func SelectionSort(list []int) {
	for i := range list {
		minIdx := i
		for j := i; j < len(list); j++ {
			if list[minIdx] > list[j] {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}

}

//BubbleSort ...
func BubbleSort(list []int) {
	for i := range list {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

//InsertionSort ...
func InsertionSort(list []int) {
	for i := range list {
		for j := i; j > 0; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}
