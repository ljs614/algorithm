package sort

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
