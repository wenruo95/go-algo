/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic150.go
*   coder: zemanzeng
*   date : 2022-03-19 22:11:43
*   desc : leetcode 150~159
*
================================================================*/

package logic

// leetcode 151: https://leetcode.com/problems/reverse-words-in-a-string/
func ReverseWords(s string) string {
	edges := make([]int, 0)

	var lastIndex int = -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			lastIndex = i
			continue
		}

		if i+1 >= len(s) || s[i+1] == ' ' {
			edges = append(edges, lastIndex+1, i+1)
		}
	}

	var result string
	for i := len(edges) - 1; i > 0; i = i - 2 {
		left, right := edges[i-1], edges[i]
		if len(result) == 0 {
			result = s[left:right]
			continue
		}
		result = result + " " + s[left:right]
	}
	return result
}

// leetcode 152: https://leetcode.com/problems/maximum-product-subarray/
func MaxProduct(nums []int) int {
	result := nums[0]
	min, max := result, result
	for i := 1; i < len(nums); i++ {
		m1, m2 := nums[i]*min, nums[i]*max
		max = maxInt(nums[i], m1, m2)
		min = minInt(nums[i], m1, m2)
		result = maxInt(result, max)
	}
	return result
}

// leetcode 153: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func FindMinItemInRotatedArray(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		if nums[left] < nums[right] {
			break
		}

		mid = (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1

		}
	}
	return nums[left]
}
