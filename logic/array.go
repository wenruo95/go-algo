/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : array.go
*   coder: zemanzeng
*   date : 2022-02-03 19:44:53
*   desc :
*
================================================================*/

package logic

// leetcode 26: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var lastIndex, dumplicate int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[lastIndex] {
			dumplicate = dumplicate + 1
			continue
		}

		nums[i-dumplicate] = nums[i]
		lastIndex = i - dumplicate
	}
	return len(nums) - dumplicate
}
