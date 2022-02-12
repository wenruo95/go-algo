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
