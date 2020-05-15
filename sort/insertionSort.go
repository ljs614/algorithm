package sort

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
