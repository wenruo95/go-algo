/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode320.go
*   coder: zemanzeng
*   date : 2022-03-19 22:02:54
*   desc : leetcode 320~329
*
================================================================*/

package logic

// leetcode 329: https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])

	list := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			list = append(list, matrix[i][j])
		}
	}

	var findLongest func(matrix [][]int, from int) int

	memo := make(map[int]int)
	findLongest = func(matrix [][]int, index int) int {
		if dist, exist := memo[index]; exist {
			return dist
		}

		// 上 下
		indexs := []int{index - n, index + n}

		row := index % n
		if row-1 >= 0 { // 左
			indexs = append(indexs, index-1)
		}
		if row+1 < n { // 右
			indexs = append(indexs, index+1)
		}

		var max int = 1
		for _, nextIndex := range indexs {
			if nextIndex >= 0 && nextIndex < len(list) &&
				list[nextIndex] > list[index] {
				if dist := findLongest(matrix, nextIndex) + 1; dist > max {
					max = dist
				}
			}
		}

		memo[index] = max
		return max
	}

	max, index := 0, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index = index + 1

			if dist := findLongest(matrix, index); dist > max {
				max = dist
			}
		}
	}
	return max
}
