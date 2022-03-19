/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic02.go
*   coder: zemanzeng
*   date : 2022-03-19 21:54:46
*   desc :
*
================================================================*/

package logic

import "container/list"

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

// leetcode 27: https://leetcode.com/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count = count + 1
		}
	}
	return count
}

// leetcode 28: https://leetcode.com/problems/implement-strstr/
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for index := 0; index < len(haystack); index++ {
		var nindex int
		for index+nindex < len(haystack) && nindex < len(needle) &&
			haystack[index+nindex] == needle[nindex] {
			nindex = nindex + 1
		}

		if nindex == len(needle) {
			return index
		}
	}
	return -1
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
