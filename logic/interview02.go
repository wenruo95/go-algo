/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : interview02.go
*   coder: zemanzeng
*   date : 2022-02-02 21:10:19
*   desc : interview sort
*
================================================================*/

package logic

import "fmt"

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

// TODO
func FindKthOfTwoSortedArray(nums1 []int, nums2 []int, k int) int {
	if k <= 0 || k > len(nums1)+len(nums2) {
		return -1
	}

	l1, r1 := 0, len(nums1)-1
	l2, r2 := 0, len(nums2)-1

	var actK int = k
	for {
		fmt.Printf("nums1(%v-%v):%v nums2(%v-%v):%v k:%v\n",
			l1, r1, nums1, l2, r2, nums2, actK)

		if r1-l1 == -1 {
			return nums2[l2+actK-1]
		}
		if r2-l2 == -1 {
			return nums1[l1+actK-1]
		}
		if actK == 1 {
			return intMin(nums1[l1], nums2[l2])
		}

		if k <= 0 {
			break
		}

		var mid1, mid2 int
		if actK/2 > r1-l1+1 {
			mid1 = r1
		} else {
			mid1 = l1 + k/2 - 1
		}
		if actK/2 > r2-l2+1 {
			mid2 = r2
		} else {
			mid2 = l2 + k/2 - 1
		}
		fmt.Printf("nums1(%v-%v-%v):%v nums2(%v-%v-%v):%v k:%v\n",
			l1, mid1, r1, nums1, l2, mid2, r2, nums2, actK)

		if nums1[mid1] > nums2[mid2] {
			actK = actK - (mid2 - l2 + 1)
			l2 = mid2 + 1
		} else {
			actK = actK - (mid1 - l1 + 1)
			l1 = mid1 + 1
		}
	}

	return -1
}
