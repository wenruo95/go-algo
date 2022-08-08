package logic

import (
	"log"
	"sort"
)

// leetcode 90: https://leetcode.com/problems/subsets-ii/
func SubsetsWithDup(nums []int) [][]int {

	var (
		track  = make([]int, len(nums))
		result = make([][]int, 0)
	)
	sort.Ints(nums)

	var backtrack func(start int, tlen int)
	backtrack = func(start int, high int) {
		result = append(result, append([]int{}, track[0:high]...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			track[high] = nums[i]
			backtrack(i+1, high+1)
		}
		log.Printf("[TEST] start:%v high:%v track:%+v result:%+v", start, high, track, jsonstr(result))
	}

	backtrack(0, 0)

	return result
}
