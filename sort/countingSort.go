package sort

//CountingSort ...
func CountingSort(list, count []int) {
	for _, val := range list {
		count[val-1]++
	}
	index := 0
	for i, val := range count {
		if val != 0 {
			for j := 0; j < val; j++ {
				list[index] = i + 1
				index++
			}
		}
	}
}
