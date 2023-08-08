package logic

// leetcode 189:https://leetcode.com/problems/rotate-array
func RotateNums(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, lo, hi int) {
	for i := 0; i < (hi-lo+1)/2; i++ {
		nums[lo+i], nums[hi-i] = nums[hi-i], nums[lo+i]
	}
}
