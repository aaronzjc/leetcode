package others

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] <= arr[j+1] {
				continue
			}
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	var qs func(left, right int)
	qs = func(left, right int) {
		if left >= right {
			return
		}
		pivot := arr[(left+right)/2]
		i, j := left, right
		for {
			for arr[i] < pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}

		qs(left, i-1)
		qs(j+1, right)
	}

	qs(0, len(arr)-1)
}

// InsertSort 插入排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		pre := i - 1
		current := arr[i]
		for pre >= 0 && arr[pre] > current {
			// 往后挪
			arr[pre+1] = arr[pre]
			pre--
		}
		arr[pre+1] = current
	}
}
