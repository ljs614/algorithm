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
