package logic

import "fmt"

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

	for st.Top() != nil {
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

func MergeSortRecur(arr []int) {
	aux := make([]int, len(arr))
	mergeSort(arr, aux, 0, len(arr)-1)
}

func mergeSort(arr, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(arr, aux, lo, mid)
	mergeSort(arr, aux, mid+1, hi)

	left1, left2 := lo, mid+1
	for index := lo; left1 <= mid || left2 <= hi; index++ {
		if left2 > hi || (left1 <= mid && arr[left1] <= arr[left2]) {
			aux[index] = arr[left1]
			left1++
			continue
		}

		aux[index] = arr[left2]
		left2++
	}

	for i := lo; i <= hi; i++ {
		arr[i] = aux[i]
	}
}

func MergeSortItera(arr []int) {
	stack := NewStack()

	auxStack := NewStack()
	auxStack.Push([]int{0, len(arr) - 1})
	auxMap := make(map[string]struct{})
	for auxStack.Top() != nil {
		value := auxStack.Pop().([]int)
		lo, hi := value[0], value[1]
		if lo >= hi {
			continue
		}

		key := fmt.Sprintf("%d_%d", lo, hi)
		if _, exist := auxMap[key]; exist {
			continue
		}
		auxMap[key] = struct{}{}

		mid := lo + (hi-lo)/2
		stack.Push([]int{lo, mid, hi})
		auxStack.Push([]int{lo, mid})
		auxStack.Push([]int{mid + 1, hi})
	}

	aux := make([]int, len(arr))
	for stack.Top() != nil {
		value := stack.Pop().([]int)
		lo, mid, hi := value[0], value[1], value[2]

		left1, left2 := lo, mid+1
		for index := lo; left1 <= mid || left2 <= hi; index++ {
			if left2 > hi || (left1 <= mid && arr[left1] <= arr[left2]) {
				aux[index] = arr[left1]
				left1++
				continue
			}
			aux[index] = arr[left2]
			left2++
		}
		for i := lo; i <= hi; i++ {
			arr[i] = aux[i]
		}
	}

}

func FindKth(arr []int, k int) int {
	return findKth(arr, len(arr)-k, 0, len(arr)-1)
}

func findKth(arr []int, k int, lo, hi int) int {
	left, right := lo, hi
	pivotValue := arr[lo]
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
	if k == pivotIndex {
		return arr[k]
	}

	if k < pivotIndex {
		return findKth(arr, k, lo, left-1)
	}
	return findKth(arr, k, left+1, hi)
}

func FindKthOfTwoSortedArray(nums1 []int, nums2 []int, k int) int {
	k = len(nums1) + len(nums2) - k + 1

	var left1, left2 int
	for left1 < len(nums1) || left2 < len(nums2) {
		if left1 == len(nums1) {
			return nums2[k-len(nums1)-1]
		}
		if left2 == len(nums2) {
			return nums1[k-len(nums2)-1]
		}

		if nums1[left1] <= nums2[left2] {
			if left1+left2+1 == k {
				return nums1[left1]
			}
			left1++
		} else {
			if left1+left2+1 == k {
				return nums2[left2]
			}
			left2++
		}
	}

	return -1
}
