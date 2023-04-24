package logic

// leetcode 371: https://leetcode.com/problems/sum-of-two-integers/
func GetSumOfTwoIntegers(a int, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		temp := sum
		sum = sum ^ carry
		carry = (carry & temp) << 1
	}
	return sum
}
