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

import (
	"strconv"
	"strings"
)

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

// leetcode 71: https://leetcode.com/problems/simplify-path/
func SimplifyPath(path string) string {
	arr, start := make([]string, 0), 0
	for i := 0; i < len(path); i++ {
		var item string
		if path[i] == '/' || i == len(path)-1 {
			if path[i] == '/' {
				item = path[start:i]
				start = i + 1
			} else {
				item = path[start : i+1]
			}
		}
		if len(item) == 0 {
			continue
		}

		switch item {
		case ".":
			continue

		case "..":
			for j := len(arr) - 1; j >= 0; j-- {
				if len(arr[j]) > 0 {
					arr[j] = ""
					break
				}
			}

		default:
			arr = append(arr, item)
		}

	}

	s := &strings.Builder{}
	s.Grow(len(path))
	for _, item := range arr {
		if len(item) == 0 {
			continue
		}
		s.WriteString("/" + item)
	}

	if s.Len() == 0 {
		return "/"
	}
	return s.String()
}

// leetcode 72: https://leetcode.com/problems/edit-distance/
// insert replace delete
func MinDistance(word1 string, word2 string) int {
	memo := make(map[string]int)

	var fn func(i1, i2 int) int
	fn = func(i1, i2 int) int {
		if i2 == len(word2) {
			return len(word1) - i1 // delete
		}
		if i1 == len(word1) {
			return len(word2) - i2 // insert
		}

		key := strconv.Itoa(i1) + "_" + strconv.Itoa(i2)
		if v, exist := memo[key]; exist {
			return v
		}

		dist1 := fn(i1, i2+1) + 1   // insert
		dist2 := fn(i1+1, i2) + 1   // delete
		dist3 := fn(i1+1, i2+1) + 1 // replace

		minDist := intMin(intMin(dist1, dist2), dist3)
		if word1[i1] == word2[i2] {
			minDist = intMin(minDist, fn(i1+1, i2+1))
		}

		memo[key] = minDist
		return minDist
	}

	return fn(0, 0)
}

// leetcode 73: https://leetcode.com/problems/set-matrix-zeroes/
func SetMatrixZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	row, column := len(matrix), len(matrix[0])
	rowWhite := make(map[int]struct{})
	columnWhite := make(map[int]struct{})

	for i := 0; i < row; i++ {
		if len(rowWhite) == row || len(columnWhite) == column {
			break
		}
		for j := 0; j < column; j++ {
			if len(columnWhite) == column {
				break
			}
			if matrix[i][j] == 0 {
				rowWhite[i] = struct{}{}
				columnWhite[j] = struct{}{}
			}
		}
	}

	for i := 0; i < row; i++ {
		if _, exist := rowWhite[i]; !exist {
			continue
		}
		for j := 0; j < column; j++ {
			matrix[i][j] = 0
		}
	}
	for i := 0; i < column; i++ {
		if _, exist := columnWhite[i]; !exist {
			continue
		}
		for j := 0; j < row; j++ {
			matrix[j][i] = 0
		}
	}
}

// leetcode 74: https://leetcode.com/problems/search-a-2d-matrix/
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, column := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[row-1][column-1] {
		return false
	}

	var x int = -1

	low, high := 0, row-1
	for low < high {
		mid := (low + high) / 2
		if matrix[mid][0] <= target && target <= matrix[mid][column-1] {
			x = mid
			break
		}
		if target > matrix[mid][0] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if matrix[low][0] <= target && target <= matrix[low][column-1] {
		x = low
	}
	if x == -1 {
		return false
	}

	low, high = 0, column-1
	for low < high {
		mid := (low + high) / 2
		if target == matrix[x][mid] {
			return true
		}
		if target > matrix[x][mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return matrix[x][low] == target
}

// leetcode 75: https://leetcode.com/problems/sort-colors/
func SortColors(nums []int) {
	var red, white, blue int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			red = red + 1
		case 1:
			white = white + 1
		case 2:
			blue = blue + 1
		default:
		}
	}

	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := red; i < white+red; i++ {
		nums[i] = 1
	}
	for i := white + red; i < len(nums); i++ {
		nums[i] = 2
	}
}

// leetcode 76: https://leetcode.com/problems/minimum-window-substring/
func MinWindow(s string, t string) string {
	getIndex := func(c byte) byte {
		var index byte
		if c >= 'a' {
			index = byte(c-'a') + 26
		} else {
			index = byte(c - 'A')
		}
		return index
	}

	set := make(map[byte]struct{})

	var arr [52]int
	for i := 0; i < len(t); i++ {
		index := getIndex(t[i])
		set[index] = struct{}{}
		arr[index] = arr[index] + 1
	}

	var low, high int
	for i := 0; i < len(s); i++ {
		if len(s)-i < len(t) {
			break
		}
		if _, exist := set[getIndex(s[i])]; !exist {
			continue
		}

		var hi int
		var arr2 [52]int
		for j := i; j < len(s); j++ {
			index := getIndex(s[j])
			if _, exist := set[index]; !exist {
				continue
			}
			if arr2[index] < arr[index] {
				arr2[index] = arr2[index] + 1
			}

			if arr2 == arr {
				hi = j + 1
				break
			}
		}

		l1, l2 := hi-i, high-low
		if (l1 > 0) && (l2 == 0 || l1 < l2) {
			low, high = i, hi
		}
	}

	return s[low:high]
}

func MinWindow2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 1. pre for t
	tvCount := make(map[byte]int)
	for index := 0; index < len(t); index++ {
		if v, exist := tvCount[t[index]]; exist {
			tvCount[t[index]] = v + 1
		} else {
			tvCount[t[index]] = 1
		}
	}

	// 2. pre for s
	svIndex := make(map[byte][]int)
	for index := 0; index < len(s); index++ {
		if v, exist := svIndex[s[index]]; exist {
			svIndex[s[index]] = append(v, index)
		} else {
			svIndex[s[index]] = []int{index}
		}
	}

	// 3. match check
	for k, count := range tvCount {
		if len(svIndex[k]) < count {
			return ""
		}
	}

	// 4. find min windows
	var low, length int = 0, len(s)
	for k, indexs := range svIndex {
		if tvCount[k] == 0 {
			continue
		}

		for j := 0; j <= len(indexs)-tvCount[k]; j++ {

			// alpham has enough count
			var right int = indexs[j]
			var invalid bool
			for k2, indexs2 := range svIndex {
				if tvCount[k2] == 0 {
					continue
				}

				var i int
				for i = 0; i < len(indexs2) && indexs2[i] < indexs[j]; i++ {
				}
				i = i + tvCount[k2] - 1
				if i >= len(indexs2) {
					invalid = true
					break
				}

				right = intMax(right, indexs2[i])
			}

			if invalid {
				break
			}

			if length > right-indexs[j]+1 {
				low = indexs[j]
				length = right - indexs[j] + 1
			}

		}

	}

	return s[low : low+length]
}
