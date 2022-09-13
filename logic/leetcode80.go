package logic

import (
	"math"
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
		columns  = make(map[int]int)
		columndp func(row, column int) int
		rows     = make(map[int]int)
		rowsdp   func(row, column int) int
	)
	columndp = func(row, column int) int {
		if row >= len(matrix) {
			return 0
		}

		key := row<<8 + column
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

		key := row<<8 + column
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

		}
	}

	return max
}

func MaximalRectangle2(matrix [][]byte) int {
	var (
		max     int
		heights = make([]int, len(matrix[0]))
	)

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == '1' {
				heights[column] = heights[column] + 1
			} else {
				heights[column] = 0
			}
		}
		max = intMax(max, LargestRectangleArea(heights))
	}
	return max
}

// leetcode 87: https://leetcode.com/problems/scramble-string/
func IsScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	var memo = make(map[string]bool)
	var scramble func(src, dst string) bool

	scramble = func(src, dst string) bool {
		if len(src) == 1 {
			return src == dst
		}
		if src == dst {
			return true
		}
		if len(src) != len(dst) {
			return false
		}

		if v, exist := memo[src+":"+dst]; exist {
			return v
		}

		var chs [26]int
		for i := 0; i < len(src); i++ {
			chs[src[i]-'a']++
			chs[dst[i]-'a']--
		}
		for _, count := range chs {
			if count != 0 {
				memo[src+":"+dst] = false
				return false
			}
		}

		var res bool
		for i := 1; i < len(src); i++ {
			res = res ||
				(scramble(src[0:i], dst[0:i]) && scramble(src[i:], dst[i:])) ||
				(scramble(src[0:i], dst[len(src)-i:]) && scramble(src[i:], dst[0:len(src)-i]))
		}
		memo[src+":"+dst] = res
		return res
	}

	return scramble(s1, s2)
}

// leetcode 88: https://leetcode.com/problems/merge-sorted-array/
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	last := len(nums1) - 1
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			last = last - 1
			m = m - 1
		} else {
			nums1[last] = nums2[n-1]
			last = last - 1
			n = n - 1
		}
	}
}

// leetcode 89: https://leetcode.com/problems/gray-code/
func GrayCode(n int) []int {
	arr := make([]int, 0, int(math.Pow(2, float64(n))))
	arr = append(arr, 0)
	for i := 0; i < n; i++ {
		length := len(arr)
		for j := length - 1; j >= 0; j-- {
			arr = append(arr, arr[j]|(1<<i))
		}
	}
	return arr
}
