/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic03.go
*   coder: zemanzeng
*   date : 2022-03-19 21:57:24
*   desc :
*
================================================================*/

package logic

import (
	"sort"
	"strconv"
	"strings"
)

// leetcode 30: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func FindSubstring(s string, words []string) []int {
	// 1. create trie tree
	type dataNode struct {
		data  map[byte]*dataNode
		item  byte
		word  string
		count int
	}

	diffCnt := 0
	root := &dataNode{data: make(map[byte]*dataNode)}
	for i := 0; i < len(words); i++ {
		var node *dataNode = root

		for j := 0; j < len(words[i]); j++ {
			b := words[i][j]

			n2, exist := node.data[b]
			if exist {
				node = n2
				continue
			}

			n2 = &dataNode{data: make(map[byte]*dataNode), item: b}
			node.data[b] = n2
			node = n2
		}

		node.word = words[i]
		node.count = node.count + 1
		if node.count == 1 {
			diffCnt = diffCnt + 1
		}
	}

	// 2. 从前往后匹配
	list := make([]int, 0)
	fullLen := len(words) * len(words[0])
	for i := 0; i < len(s); i++ {
		var node *dataNode = root

		matchCnt := 0
		set := make(map[string]int)
		for j := 0; j < fullLen && i+j < len(s); j++ {
			n2, exist := node.data[s[i+j]]
			if !exist {
				break
			}

			if len(n2.word) > 0 { // 匹配到叶子节点
				if set[n2.word] >= n2.count { // 有重复匹配问题
					break
				}
				set[n2.word] = set[n2.word] + 1
				matchCnt = matchCnt + 1

				node = root
				continue
			}

			node = n2
		}

		if len(set) == diffCnt && matchCnt == len(words) {
			list = append(list, i)
			continue
		}
	}

	return list
}

func FindSubstring2(s string, words []string) []int {
	if len(words) == 0 || len(words)*len(words[0]) > len(s) {
		return nil
	}

	wordCnt := make(map[string]int)
	for _, word := range words {
		wordCnt[word] = wordCnt[word] + 1
	}

	size := len(words[0])
	fullLen := len(words) * size

	list := make([]int, 0)
	for i := 0; i < len(s)-fullLen+1; i++ {

		sum := 0
		matchCnt := make(map[string]int)
		for j := 0; i+j+size <= len(s); j = j + size {
			word := s[i+j : i+j+size]
			if _, exist := wordCnt[word]; !exist {
				break
			}
			if matchCnt[word] == wordCnt[word] {
				break
			}

			matchCnt[word] = matchCnt[word] + 1
			sum = sum + 1
		}

		if sum == len(words) && len(matchCnt) == len(wordCnt) {
			list = append(list, i)
		}

	}

	return list
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

// leetcode 31: https://leetcode.com/problems/next-permutation/
// [1,2,3] => [1,3,2] => [3,1,2] => [2,3,1] => [3,1,2] -> [3,2,1]
func NextPermutation(nums []int) {
	var k int
	for k = len(nums) - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	if k < 0 {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
		return
	}

	var l int
	for l = len(nums) - 1; l > k; l-- {
		if nums[k] < nums[l] {
			break
		}
	}
	nums[k], nums[l] = nums[l], nums[k]

	// [k + 1:]
	start := k + 1
	for i := 0; i < (len(nums)-start)/2; i++ {
		left, right := start+i, len(nums)-1-i
		nums[left], nums[right] = nums[right], nums[left]
	}

}

// leetcode 32: https://leetcode.com/problems/longest-valid-parentheses/
func LongestValidParentheses(s string) int {
	var maxlen int
	for index := 0; index < len(s); index++ {
		var left, right int
		for j := index; s[index] == '(' && j < len(s); j++ {
			if s[j] == '(' {
				left = left + 1
				continue
			}

			right = right + 1
			if left == right {
				maxlen = intmax(maxlen, left+right)
				continue
			}

			if left < right {
				index = j - 1
				maxlen = intmax(maxlen, left+right-1)
				break
			}
		}

	}

	return maxlen
}

func intmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode 34: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func SearchRange(nums []int, target int) []int {
	var index int = -1

	// 1. find mid
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			index = mid
			break
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	result := make([]int, 0)

	// 2. find left
	left, right := low, index
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid-1 < low || nums[mid-1] < target {
				result = append(result, mid)
				break
			}
			right = mid - 1
			continue
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
	if len(result) < 1 {
		result = append(result, index)
	}

	// 3. find right
	left, right = index, high
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid+1 > high || nums[mid+1] > target {
				result = append(result, mid)
				break
			}
			left = mid + 1
			continue
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}
	if len(result) < 2 {
		result = append(result, index)
	}

	return result
}

// leetcode 35: https://leetcode.com/problems/search-insert-position/
func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1

	var mid int
	for low <= high {
		mid = (low + high) / 2

		if nums[mid] >= target {
			if mid-1 < 0 || nums[mid-1] < target {
				return mid
			}

			high = mid - 1
			continue
		}

		low = mid + 1
	}

	return mid + 1
}

// leetcode 36: https://leetcode.com/problems/valid-sudoku/
func IsValidSudoku(board [][]byte) bool {
	if len(board) == 0 {
		return false
	}

	columns, rows := len(board), len(board[0])
	// column
	for column := 0; column < columns; column++ {
		set := make(map[byte]struct{})
		for row := 0; row < rows; row++ {
			b := board[column][row]
			if _, exist := set[b]; exist {
				return false
			}
			if b >= '0' && b <= '9' {
				set[b] = struct{}{}
			}
		}
	}

	// row
	for row := 0; row < rows; row++ {
		set := make(map[byte]struct{})
		for column := 0; column < columns; column++ {
			b := board[column][row]
			if _, exist := set[b]; exist {
				return false
			}
			if b >= '0' && b <= '9' {
				set[b] = struct{}{}
			}
		}
	}

	// 3*3
	for column := 0; column+2 < columns; column = column + 3 {
		for row := 0; row+2 < rows; row = row + 3 {
			set := make(map[byte]struct{})

			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					b := board[column+j][row+k]
					if _, exist := set[b]; exist {
						return false
					}
					if b >= '0' && b <= '9' {
						set[b] = struct{}{}
					}
				}
			}

		}
	}

	return true
}

// leetcode 37: https://leetcode.com/problems/sudoku-solver/
func SolveSudoku(board [][]byte) {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}

	var solve func(board [][]byte) bool

	solve = func(board [][]byte) bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {

				if board[i][j] == '.' {
					for c := byte('1'); c <= '9'; c++ {
						if isValidChar(board, i, j, c) {
							board[i][j] = c

							if solve(board) {
								return true
							} else {
								board[i][j] = '.'
							}

						}
					}

					return false
				}

			}
		}

		return true
	}

	solve(board)
}

func isValidChar(board [][]byte, row int, column int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == c || board[i][column] == c {
			return false
		}
		if x, y := (row/3)*3+i/3, (column/3)*3+i%3; board[x][y] == c {
			return false
		}
	}
	return true
}

// leetcode 38: https://leetcode.com/problems/count-and-say/
func CountAndSay(n int) string {

	var s string = "1"

	for j := 2; j <= n; j++ {
		result := &strings.Builder{}

		var count int
		for i := 0; i < len(s); i++ {
			count = count + 1

			if i+1 < len(s) && s[i] == s[i+1] {
				continue
			}

			result.WriteString(strconv.Itoa(count))
			result.WriteByte(s[i])
			count = 0
		}

		s = result.String()
	}

	return s
}

// leetcode 39: https://leetcode.com/problems/combination-sum/
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)

	memo := make(map[int][][]int)

	var fn func(target int) [][]int
	fn = func(target int) [][]int {
		if target < candidates[0] {
			return nil
		}
		if v, exist := memo[target]; exist {
			return v
		}

		arrays := make([][]int, 0)
		for i := 0; i < len(candidates); i++ {
			if target == candidates[i] {
				arrays = append(arrays, []int{candidates[i]})
				continue
			}

			lists := fn(target - candidates[i])
			for _, list := range lists {
				if len(list) > 0 && candidates[i] > list[0] { // 确保数据有序
					continue
				}
				arrays = append(arrays,
					append([]int{candidates[i]}, list...),
				)
			}
		}

		memo[target] = arrays
		return arrays
	}

	return fn(target)
}
