/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode70.go
*   coder: zemanzeng
*   date : 2022-04-03 22:01:46
*   desc : leetcode 70~79
*
================================================================*/

package logic

// leetcode 70: https://leetcode.com/problems/climbing-stairs/
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	n1, n2, x := 1, 2, 2
	for x < n {
		x = x + 1
		n1, n2 = n2, n1+n2
	}
	return n2
}
