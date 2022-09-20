package logic

// leetcode 191: https://leetcode.com/problems/number-of-1-bits/
func HammingWeight(num uint32) int {
	var count int
	for num > 0 {
		count = count + 1
		num = num & (num - 1)
	}
	return count
}
