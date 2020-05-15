package sort

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
