/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode60.go
*   coder: zemanzeng
*   date : 2022-03-29 10:24:06
*   desc : leetcode 60~69
*
================================================================*/

package logic

import (
	"strconv"
	"strings"
)

// leetcode 60: https://leetcode.com/problems/permutation-sequence/
// 1 {2, 3, 4}
// 2 {1, 3, 4}
// 3 {1, 2, 4}
// 4 {1, 2, 3}
func GetPermutation(n int, k int) string {
	factorial := make([]int, n) // 1, 1, 2, 6, 24
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = i * factorial[i-1]
	}

	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	s := &strings.Builder{}
	s.Grow(k)

	k = k - 1
	for i := 1; i <= n; i++ {
		index := k / factorial[n-i]
		s.WriteString(strconv.Itoa(nums[index]))
		nums = append(nums[0:index], nums[index+1:]...)
		k = k - index*factorial[n-i]
	}

	return s.String()
}

// leetcode 62: https://leetcode.com/problems/unique-paths/
// 本质上为排列组合(m-1)个A和(n-1)个B
func UniquePaths(m int, n int) int {

	step := (m - 1) + (n - 1)
	pivot := intMin(m-1, n-1)

	var unique int = 1
	var repeated int = 1
	for i := 0; i < pivot; i++ {
		unique = unique * (step - i)
		repeated = repeated * (i + 1)
		if unique%repeated == 0 {
			unique = unique / repeated
			repeated = 1
		}
	}
	return unique / repeated
}
