/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode50.go
*   coder: zemanzeng
*   date : 2022-03-24 21:59:14
*   desc : leetcode 50~59
*
================================================================*/

package logic

import (
	"fmt"
	"strconv"
)

// leetcode 50: https://leetcode.com/problems/powx-n/
func MyPow(x float64, n int) float64 {
	memo := make(map[int]float64)

	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		if n == 1 {
			return x
		}
		if n < 0 {
			return 1 / pow(x, -n)
		}
		if v, exist := memo[n]; exist {
			return v
		}

		result := pow(x, n/2) * pow(x, n/2) * pow(x, n-(n/2)*2)
		memo[n] = result
		return result
	}

	result := pow(x, n)
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", result), 64) // 保留5为小数 规避9.261000000000001问题
	return v
}

// leetcode 51: https://leetcode.com/problems/n-queens/
func SolveNQueens(n int) [][]string {
	array := make([][]byte, n)
	for i := 0; i < n; i++ {
		bts := make([]byte, n)
		for j := 0; j < n; j++ {
			bts[j] = '.'
		}
		array[i] = bts
	}

	var queens func(array [][]byte, n int, row int)
	var check func(array [][]byte, n int, row int, column int) bool

	arrays := make([][]string, 0)
	queens = func(array [][]byte, n int, row int) {
		if row >= n {
			list := make([]string, 0)
			for i := 0; i < len(array); i++ {
				list = append(list, string(array[i]))
			}
			arrays = append(arrays, list)
			return
		}

		for y := 0; y < n; y++ {
			if check(array, n, row, y) {
				array[row][y] = 'Q'
				queens(array, n, row+1)
				array[row][y] = '.'
			}
		}
	}

	check = func(array [][]byte, n int, row int, column int) bool {
		for x := row - 1; x >= 0; x-- {
			if array[x][column] == 'Q' {
				return false
			}
			if slope := column - (row - x); slope >= 0 && array[x][slope] == 'Q' {
				return false
			}
			if slope := column + (row - x); slope < n && array[x][slope] == 'Q' {
				return false
			}
		}
		return true
	}

	queens(array, n, 0)
	return arrays
}

// leetcode 52: https://leetcode.com/problems/n-queens-ii/
func TotalNQueens(n int) int {

	array := make([][]byte, n)
	for i := 0; i < n; i++ {
		bts := make([]byte, n)
		for j := 0; j < n; j++ {
			bts[j] = '.'
		}
		array[i] = bts
	}

	var queens func(array [][]byte, n int, row int)
	var check func(array [][]byte, n int, row int, column int) bool

	var cnt int
	queens = func(array [][]byte, n int, row int) {
		if row >= n {
			cnt = cnt + 1
			return
		}

		for y := 0; y < n; y++ {
			if check(array, n, row, y) {
				array[row][y] = 'Q'
				queens(array, n, row+1)
				array[row][y] = '.'
			}
		}
	}

	check = func(array [][]byte, n int, row int, column int) bool {
		for x := row - 1; x >= 0; x-- {
			if array[x][column] == 'Q' {
				return false
			}
			if slope := column - (row - x); slope >= 0 && array[x][slope] == 'Q' {
				return false
			}
			if slope := column + (row - x); slope < n && array[x][slope] == 'Q' {
				return false
			}
		}
		return true
	}

	queens(array, n, 0)
	return cnt
}

// leetcode 53: https://leetcode.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max int = nums[0]
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > max {
				max = sum
			}

		}
	}
	return max
}

func MaxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + intMax(dp[i-1], 0)
		max = intMax(max, dp[i])
	}
	return max
}

// leetcode 54: https://leetcode.com/problems/spiral-matrix/
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	row, column := len(matrix), len(matrix[0])

	cnt := intMin(row, column) / 2
	if intMin(row, column)%2 == 1 {
		cnt = cnt + 1
	}

	list := make([]int, 0)
	for i := 0; i < cnt; i++ { // 最外层
		if i == row-i-1 && i == column-i-1 { // 中间只有一个元素的情况
			list = append(list, matrix[i][i])
			continue
		}

		// 右
		for y := i; y <= column-i-1; y++ {
			list = append(list, matrix[i][y])
		}

		// 下
		if i == row-i-1 {
			break
		}
		for x := i + 1; x < row-i-1; x++ {
			list = append(list, matrix[x][column-i-1])
		}

		// 左
		for y := column - i - 1; y >= i; y-- {
			list = append(list, matrix[row-i-1][y])
		}

		// 上
		if i == column-i-1 {
			break
		}
		for x := row - i - 2; x > i; x-- {
			list = append(list, matrix[x][i])
		}

	}

	return list
}
