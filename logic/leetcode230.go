package logic

// leetcode 238: https://leetcode.com/problems/product-of-array-except-self/
func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	rightValue := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * rightValue
		rightValue = rightValue * nums[i]
	}
	return result
}
