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
