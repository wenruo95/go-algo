/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic.go
*   coder: zemanzeng
*   date : 2022-02-03 00:49:06
*   desc :
*
================================================================*/

package logic

import (
	"container/list"
)

// leetcode 20:https://leetcode.com/problems/valid-parentheses/
func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	left := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			left.PushBack(s[i])
			continue
		}
		if left.Len() != 0 && pairs[s[i]] == left.Back().Value.(byte) {
			left.Remove(left.Back())
			continue
		}
		return false
	}

	return left.Len() == 0
}

// leetcode 22: https://leetcode.com/problems/generate-parentheses/
func GenerateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	memo := make(map[int][]string)
	memo[1] = []string{"()"}

	var parenthesis func(n int) []string
	parenthesis = func(n int) []string {
		if list, exist := memo[n]; exist {
			return list
		}

		// f(n - 1) + f(1)
		results := make(map[string]struct{})
		for _, s := range parenthesis(n - 1) {
			results["("+s+")"] = struct{}{}
			results["()"+s] = struct{}{}
			results[s+"()"] = struct{}{}
		}

		// f(n-x) + f(x)
		left, right := 2, n-2
		for left <= right {
			list1 := parenthesis(left)
			list2 := parenthesis(right)
			for _, item1 := range list1 {
				for _, item2 := range list2 {
					results[item1+item2] = struct{}{}
					results[item2+item1] = struct{}{}
				}
			}
			left = left + 1
			right = right - 1
		}

		list := make([]string, 0)
		for s := range results {
			list = append(list, s)
		}

		memo[n] = list
		return list
	}

	return parenthesis(n)
}

// leetcode 29: https://leetcode.com/problems/divide-two-integers/
func Divide(dividend int, divisor int) int {

	positive := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if divisor == 0 || dividend == 0 || dividend < divisor {
		return 0
	}

	var bitcnt int
	for (divisor << (bitcnt + 1)) <= dividend {
		bitcnt = bitcnt + 1
	}

	cnt := 0
	left := dividend - (divisor << bitcnt)
	for left-divisor >= 0 {
		cnt = cnt + 1
		left = left - divisor
	}

	res := (1 << bitcnt) + cnt
	if !positive {
		res = -res
	}

	if max := (1<<31 - 1); res > max {
		return max
	}
	if min := -(1 << 31); res < min {
		return min
	}
	return res
}

// leetcode 33: https://leetcode.com/problems/search-in-rotated-sorted-array/
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
