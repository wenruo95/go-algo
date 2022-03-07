/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sort.go
*   coder: zemanzeng
*   date : 2022-02-02 21:10:19
*   desc : sort
*
================================================================*/

package logic

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	a, b := left, right
	pivotValue := nums[left]
	for a < b {
		for ; b > a && nums[b] >= pivotValue; b-- {
		}
		for ; a < b && nums[a] <= pivotValue; a++ {
		}

		nums[b], nums[a] = nums[a], nums[b]
	}
	nums[left] = nums[a]
	nums[a] = pivotValue

	pivot := b
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func FindKth(a []int, k int) int {
	return findKth(a, len(a)-k, 0, len(a)-1)
}

func findKth(nums []int, k int, low int, high int) int {
	parti := nums[low]

	left, right := low, high
	for left < right {
		for left < right && nums[right] >= parti {
			right = right - 1
		}
		if left == right {
			break
		}

		nums[left] = nums[right]
		left = left + 1

		for left < right && nums[left] <= parti {
			left = left + 1
		}
		nums[right] = nums[left]
		right = right - 1
	}
	nums[left] = parti

	if left < k {
		return findKth(nums, k, left+1, high)
	} else if left > k {
		return findKth(nums, k, low, left-1)
	}

	return nums[k]
}
