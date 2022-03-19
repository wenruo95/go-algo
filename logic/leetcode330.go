/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode330.go
*   coder: zemanzeng
*   date : 2022-03-19 22:02:03
*   desc : leetcode 330~339
*
================================================================*/

package logic

// leetcode 338: https://leetcode.com/problems/counting-bits/submissions/
func CountBits(n int) []int {
	if n < 0 {
		return nil
	}

	// n * log(n)
	list := make([]int, n+1)
	for i := 0; i <= n; i++ {
		sum, num := 0, i
		for num != 0 {
			sum = sum + num%2
			num = num / 2
		}
		list[i] = sum
	}

	return list
}
