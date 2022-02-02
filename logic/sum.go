/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sum.go
*   coder: zemanzeng
*   date : 2022-02-02 21:01:05
*   desc : sum
*
================================================================*/

package logic

import "sort"

// leetcode 16: https://leetcode.com/problems/3sum-closest/
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var mindist int
	var indexs [3]int
	for i := 0; i < len(nums)-2; i++ {
		i1, i2 := twoSumClosest(nums, i+1, target-nums[i])
		if distance := abs(nums[i]+nums[i1]+nums[i2], target); i == 0 || distance < mindist {
			mindist = distance
			indexs[0], indexs[1], indexs[2] = i, i1, i2
		}
	}

	return nums[indexs[0]] + nums[indexs[1]] + nums[indexs[2]]
}

func abs(num, target int) int {
	if num > target {
		return num - target
	}
	return target - num
}

func twoSumClosest(nums []int, start int, target int) (int, int) {
	low, high := start, len(nums)-1

	i1, i2 := low, high
	min := abs(nums[low]+nums[high], target)
	for low < high {
		sum := nums[low] + nums[high]
		if distance := abs(sum, target); distance < min {
			min = distance
			i1, i2 = low, high
		}

		if sum > target {
			high = high - 1
		} else if sum == target {
			break
		} else {
			low = low + 1
		}
	}
	return i1, i2
}
