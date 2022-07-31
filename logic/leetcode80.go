package logic

import (
	"fmt"
	"strconv"
)

// leetcode 80: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func RemoveDuplicates2(nums []int) int {
	var count, k int

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			count = 0
		}

		count = count + 1
		if count > 2 {
			k = k + 1
		}

		nums[i-k] = nums[i]
	}
	return len(nums) - k
}

// leetcode 81: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func SearchInRotatedSortedArrayII(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {

		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left = left + 1
			right = right - 1
		} else if nums[left] <= nums[mid] {
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

	return false
}

// leetcode 84: https://leetcode.com/problems/largest-rectangle-in-histogram/
func LargestRectangleArea(heights []int) int {
	var (
		area  = 0
		stack = NewStack()
	)
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for !stack.Empty() && heights[stack.Top()] >= heights[i] {
			cur := stack.Top()
			stack.Pop()

			width := i
			if !stack.Empty() {
				width = (i - stack.Top() - 1)
			}
			area = intMax(area, heights[cur]*width)
		}
		stack.Push(i)
	}
	return area
}

// leetcode 85: https://leetcode.com/problems/maximal-rectangle/
func MaximalRectangle(matrix [][]byte) int {
	var (
		columns  = make(map[string]int)
		columndp func(row, column int) int
		rows     = make(map[string]int)
		rowsdp   func(row, column int) int
	)
	columndp = func(row, column int) int {
		if row >= len(matrix) {
			return 0
		}

		key := strconv.Itoa(row) + "_" + strconv.Itoa(column)
		if v, exist := columns[key]; exist {
			return v
		}

		var result int
		if matrix[row][column] == '1' {
			result = 1 + columndp(row+1, column)
		} else {
			result = 0
		}
		columns[key] = result
		return result
	}
	rowsdp = func(row, column int) int {
		if column >= len(matrix[0]) {
			return 0
		}

		key := strconv.Itoa(row) + "_" + strconv.Itoa(column)
		if v, exist := rows[key]; exist {
			return v
		}

		var result int
		if matrix[row][column] == '1' {
			result = 1 + rowsdp(row, column+1)
		} else {
			result = 0
		}
		rows[key] = result
		return result
	}

	var max int
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			maxRow := rowsdp(row, column)

			var minColumn int = len(matrix) + 1
			for width := 1; width <= maxRow && matrix[row][column+width-1] != 0; width++ {
				minColumn = intMin(minColumn, columndp(row, column+width-1))
				max = intMax(max, width*minColumn)
			}
			fmt.Printf("[TEST0] row:%v column:%v max_row:%v min_column:%v max:%v\n",
				row, column, maxRow, minColumn, max)

		}
	}
	fmt.Printf("[TEST9] rows:%v columns:%v max:%v\n", jsonstr(rows), jsonstr(columns), max)

	return max
}
