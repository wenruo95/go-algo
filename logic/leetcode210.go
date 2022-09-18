package logic

// leetcode 217: https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, exist := set[nums[i]]; exist {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}
