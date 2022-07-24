package logic

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
