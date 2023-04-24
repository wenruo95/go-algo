package logic

// Rec: recursion
// Iter: iteration

func QuickSortRecur(arr []int) {
	quickSortRecur(arr, 0, len(arr)-1)
	return
}

func quickSortRecur(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	left, right := lo, hi
	pivotValue := arr[left]
	for left < right {
		for left < right && arr[right] >= pivotValue {
			right--
		}
		for left < right && arr[left] <= pivotValue {
			left++
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[lo] = arr[left]
	arr[left] = pivotValue

	pivotIndex := left
	quickSortRecur(arr, lo, pivotIndex-1)
	quickSortRecur(arr, pivotIndex+1, hi)
}

func QuickSortItera(arr []int) {
	st := NewStack()
	st.Push([]int{0, len(arr) - 1})
	for st.Size() > 0 {
		value := st.Pop().([]int)
		lo, hi := value[0], value[1]
		if lo >= hi {
			continue
		}

		left, right := lo, hi
		pivotValue := arr[left]
		for left < right {
			for left < right && arr[right] >= pivotValue {
				right--
			}
			for left < right && arr[left] <= pivotValue {
				left++
			}
			arr[left], arr[right] = arr[right], arr[left]
		}
		arr[lo] = arr[left]
		arr[left] = pivotValue

		pivotIndex := left
		st.Push([]int{lo, pivotIndex - 1})
		st.Push([]int{pivotIndex + 1, hi})
	}

}
