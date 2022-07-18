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
