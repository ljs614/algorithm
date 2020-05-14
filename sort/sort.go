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

//QuickSort ...
func QuickSort(list []int, start, end int) {
	if start >= end {
		return
	}
	pivot := start // pivot은 첫번째 원소
	i := start + 1
	j := end
	for i <= j { // 엇갈릴 때까지 반복
		for i <= end && list[i] <= list[pivot] { // pivot보다 큰 값을 만날 때까지 왼쪽에서부터
			i++
		}
		for j > start && list[j] >= list[pivot] { // pivot보다 작은 값을 만날 때까지 오른쪽에서부터
			j--
		}
		if i > j { // 엇갈린 상태이면 pivot과 큰값을 교체
			list[pivot], list[j] = list[j], list[pivot]
		} else { // 엇갈리지 않았으면 작은 값과 큰 값을 교체
			list[i], list[j] = list[j], list[i]
		}
	}

	// pivot을 기준으로 배열을 두개로 나눈 후 재귀
	QuickSort(list, start, j-1)
	QuickSort(list, j+1, end)
}

//QuickSortDesc ...
func QuickSortDesc(list []int, start, end int) {
	if start >= end {
		return
	}
	pivot := start // pivot은 첫번째 원소
	i := start + 1
	j := end
	for i <= j { // 엇갈릴 때까지 반복
		for i <= end && list[i] >= list[pivot] { // pivot보다 작은 값을 만날 때까지 왼쪽에서부터
			i++
		}
		for j > start && list[j] <= list[pivot] { // pivot보다 큰 값을 만날 때까지 오른쪽에서부터
			j--
		}
		if i > j { // 엇갈린 상태이면 pivot과 작은값을 교체
			list[pivot], list[j] = list[j], list[pivot]
		} else { // 엇갈리지 않았으면 작은 값과 큰 값을 교체
			list[i], list[j] = list[j], list[i]
		}
	}

	// pivot을 기준으로 배열을 두개로 나눈 후 재귀
	QuickSortDesc(list, start, j-1)
	QuickSortDesc(list, j+1, end)
}
