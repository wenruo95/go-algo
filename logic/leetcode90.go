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

// https://leetcode.com/problems/decode-ways/
func NumDecodings(s string) int {

	var (
		dp   func(s string, start int) int
		memo = make(map[int]int)
	)

	dp = func(s string, start int) int {
		if start > len(s)-1 {
			return 1
		}
		if s[start] <= '0' || s[start] > '9' {
			return 0
		}
		if v, exist := memo[start]; exist {
			return v
		}

		count := dp(s, start+1)
		if start+1 < len(s) {
			if sum := (s[start]-'0')*10 + (s[start+1] - '0'); sum > 0 && sum <= 26 {
				count = count + dp(s, start+2)
			}
		}

		memo[start] = count
		return count
	}

	return dp(s, 0)
}
