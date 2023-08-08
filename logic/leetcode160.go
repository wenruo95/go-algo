package logic

import "sort"

// leetcode 169: https://leetcode.com/problems/majority-element
func MajorityElement(nums []int) int {
	sort.Ints(nums)

	step := len(nums) / 2
	if len(nums)%2 == 1 {
		step = step + 1
	}

	for i := 0; i < len(nums); i++ {
		if i+step-1 < len(nums) && nums[i] == nums[i+step-1] {
			return nums[i]
		}
		for j := i; j < len(nums); j = j + 2 {
			if nums[i] == nums[j] {
				continue
			}
			i = j - 2
			break
		}
	}
	return nums[0]
}
