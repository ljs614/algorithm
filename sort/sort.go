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

//MergeSort ...
func MergeSort(list []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSort(list, start, mid)
	MergeSort(list, mid+1, end)
	merge(list, start, mid, end)
}

func merge(list []int, start, mid, end int) {
	sorted := make([]int, len(list))
	i := start
	j := mid + 1
	k := start
	for i <= mid || j <= end {
		if i > mid {
			sorted[k] = list[j]
			j++
		} else if j > end {
			sorted[k] = list[i]
			i++
		} else if list[i] <= list[j] {
			sorted[k] = list[i]
			i++
		} else {
			sorted[k] = list[j]
			j++
		}
		k++
	}

	for x := start; x <= end; x++ {
		list[x] = sorted[x]
	}
}

//HeapSort ...
func HeapSort(list []int) {
	len := len(list)

	//힙을 구성
	for i := 1; i < len; i++ {
		c := i
		for {
			root := (c - 1) / 2
			if list[root] < list[c] {
				list[root], list[c] = list[c], list[root]
			}
			if root == 0 {
				break
			}
			c = root
		}
	}

	// 크기를 줄여가며 반복적으로 힙을 구성
	for i := len - 1; i >= 0; i-- {
		list[0], list[i] = list[i], list[0]
		root := 0
		c := 1
		for {
			c = 2*root + 1
			//자식중 더 큰 값 찾기
			if c < i-1 && list[c] < list[c+1] {
				c++
			}
			// 루트보다 자식이 크면 교환
			if c < i && list[root] < list[c] {
				list[root], list[c] = list[c], list[root]
			}
			root = c
			if c >= i {
				break
			}
		}
	}
}

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
