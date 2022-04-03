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

// leetcode 63: https://leetcode.com/problems/unique-paths-ii/
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	var fn func(x, y int) int

	row, column := len(obstacleGrid), len(obstacleGrid[0])
	memo := make(map[int]int)
	fn = func(x, y int) int {
		if x >= row || y >= column || obstacleGrid[x][y] == 1 {
			return 0
		}
		if x == row-1 && y == column-1 {
			return 1
		}
		if v, exist := memo[x*column+y]; exist {
			return v
		}

		count := fn(x+1, y) + fn(x, y+1)
		memo[x*column+y] = count
		return count
	}

	return fn(0, 0)
}

// leetcode 64: https://leetcode.com/problems/minimum-path-sum/
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var fn func(x, y int) int

	memo := make(map[int]int)
	row, column := len(grid), len(grid[0])
	fn = func(x, y int) int {
		if x >= row || y >= column {
			return -1
		}
		if x == row-1 && y == column-1 {
			return grid[x][y]
		}

		key := (x * column) + y
		if v, exist := memo[key]; exist {
			return v
		}

		right := fn(x, y+1) // 右
		down := fn(x+1, y)  // 下

		var sum int
		if right == -1 || down == -1 {
			sum = intMax(right, down) + grid[x][y]
		} else {
			sum = intMin(right, down) + grid[x][y]
		}

		memo[key] = sum
		return sum
	}

	return fn(0, 0)
}

// leetcode 65: https://leetcode.com/problems/valid-number/
func IsNumber(s string) bool {
	// + -  第一位, 后面要接数字
	// .	只能够出现一次 且在e之前,不能够单独出现
	// e	前后均要接数字, 只能够出现一次
	// 0 1 2 3 4 5 6 7 8 9

	point, echar, number := -1, -1, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-':
			if i == 0 {
				continue
			}
			return false

		case '.':
			if point == -1 && echar == -1 {
				point = i
				continue
			}
			return false

		case 'e', 'E':
			if number == 0 {
				return false
			}
			if echar == -1 && i-1 >= 0 && i+1 < len(s) {
				echar = i

				if (s[i+1] == '+' || s[i+1] == '-') && i+2 < len(s) {
					i = i + 1
				}
				continue
			}
			return false

		default:
			if s[i] >= '0' && s[i] <= '9' {
				number = number + 1
				continue
			}
			return false
		}
	}

	return number > 0
}

// leetcode 66: https://leetcode.com/problems/plus-one/
func PlusOne(digits []int) []int {
	var higher int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if higher == 0 {
			break
		}
		sum := higher + digits[i]
		higher = sum / 10
		digits[i] = sum % 10
	}

	if higher > 0 {
		return append([]int{higher}, digits...)
	}
	return digits
}

// leetcode 67: https://leetcode.com/problems/add-binary/
func AddBinary(a string, b string) string {
	result := make([]byte, 1+intMax(len(a), len(b)))
	i, j, k := len(a)-1, len(b)-1, len(result)-1

	var higher byte
	for (i >= 0 && j >= 0) || higher > 0 {
		sum := higher
		if i >= 0 {
			sum = sum + a[i] - '0'
		}
		if j >= 0 {
			sum = sum + b[j] - '0'
		}
		higher = sum / 2
		result[k] = sum%2 + '0'

		i = i - 1
		j = j - 1
		k = k - 1
	}

	if i >= 0 {
		return a[0:i+1] + string(result[k+1:])
	}
	if j >= 0 {
		return b[0:j+1] + string(result[k+1:])
	}
	return string(result[k+1:])
}

// leetcode 68: https://leetcode.com/problems/text-justification/
func FullJustify(words []string, maxWidth int) []string {
	list := make([]string, 0)

	var start, sum int
	for i := 0; i < len(words); i++ {
		width := sum + i - start + len(words[i])
		if width >= maxWidth {
			list = append(list, strings.Join(words[start:i+1], " "))
			start, sum = i+1, 0
			continue
		}

		sum = sum + len(words[i])
		if i+1 < len(words) && width+1+len(words[i+1]) <= maxWidth {
			continue
		}

		if i == start || i == len(words)-1 {
			list = append(list, strings.Join(words[start:i+1], " ")+
				strings.Repeat(" ", maxWidth-width))
			start, sum = i+1, 0
			continue
		}

		space := maxWidth - sum
		avg := space / (i - start)
		left := space - (i-start)*avg
		holder := strings.Repeat(" ", avg)

		b := &strings.Builder{}
		b.Grow(maxWidth)
		for j := start; j < i; j++ {
			b.WriteString(words[j] + holder)

			if left > 0 {
				left = left - 1
				b.WriteString(" ")
			}
		}
		b.WriteString(words[i])

		list = append(list, b.String())
		start, sum = i+1, 0
	}

	return list
}
