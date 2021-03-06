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
	"sort"
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

// leetcode 55: https://leetcode.com/problems/jump-game/
func CanJump(nums []int) bool {
	var farthest int
	for i := 0; i < len(nums)-1; i++ {
		if farthest >= i {
			farthest = intMax(farthest, i+nums[i])
		}
	}
	return farthest >= len(nums)-1
}

// leetcode 56: https://leetcode.com/problems/merge-intervals/
func MergeIntervals(intervals [][]int) [][]int {
	left2Index := make(map[int][]int)
	for index, collect := range intervals {
		if v, exist := left2Index[collect[0]]; exist {
			left2Index[collect[0]] = append(v, index)
		} else {
			left2Index[collect[0]] = []int{index}
		}
	}

	lefts := make([]int, 0)
	for left := range left2Index {
		lefts = append(lefts, left)
	}
	sort.Ints(lefts)

	lists := make([][]int, 0)

	var low, high int = -1, -1
	for index, left := range lefts {
		if low == -1 {
			low = left
		}

		for _, index := range left2Index[left] {
			high = intMax(intervals[index][1], high)
		}

		if index >= len(lefts)-1 || high < lefts[index+1] {
			lists = append(lists, []int{low, high})
			low = -1
		}

	}

	return lists
}

// leetcode 57: https://leetcode.com/problems/insert-interval/
func MergeInsertInterval(intervals [][]int, newInterval []int) [][]int {
	left, right := -1, -1
	for index := 0; index < len(intervals); index++ {
		if left == -1 && newInterval[0] <= intervals[index][1] {
			left = index
		}
		if right == -1 && newInterval[1] < intervals[index][0] {
			right = index
		}
	}
	if left == -1 { // [0, len-1] [left...]
		return append(intervals, newInterval)
	}
	if right == -1 {
		right = len(intervals)
	}

	minLeft := intMin(intervals[left][0], newInterval[0]) // 左区间
	maxRight := newInterval[1]
	if right-1 >= 0 {
		maxRight = intMax(intervals[right-1][1], newInterval[1]) // 右区间
	}
	if right+1 <= len(intervals)-1 && intervals[right+1][1] == maxRight { // 判断是否与后一个合并
		right = right + 1
		maxRight = intervals[right][1]
	}

	arrays := make([][]int, 0)
	arrays = append(arrays, intervals[0:left]...)
	arrays = append(arrays, []int{minLeft, maxRight})
	arrays = append(arrays, intervals[right:]...)
	return arrays
}

// leetcode 58: https://leetcode.com/problems/length-of-last-word/
func LengthOfLastWord(s string) int {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if s[left] == ' ' || (i-1 >= 0 && s[i-1] == ' ') {
			left = i
		}
		right = i
	}

	return right - left + 1
}

// leetcode 59: https://leetcode.com/problems/spiral-matrix-ii/
func GenerateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	var seq int = 0
	for i := 0; i <= n/2; i++ {
		// 右
		for row := i; row <= n-i-1; row++ {
			seq = seq + 1
			matrix[i][row] = seq
		}
		if seq == n*n {
			break
		}

		// 下
		for column := i + 1; column <= n-i-2; column++ {
			seq = seq + 1
			matrix[column][n-i-1] = seq
		}

		// 左
		for row := n - i - 1; row >= i; row-- {
			seq = seq + 1
			matrix[n-i-1][row] = seq
		}
		if seq == n*n {
			break
		}

		// 上
		for column := n - i - 2; column > i; column-- {
			seq = seq + 1
			matrix[column][i] = seq
		}
	}

	return matrix
}
