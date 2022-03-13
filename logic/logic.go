/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic.go
*   coder: zemanzeng
*   date : 2022-02-03 00:49:06
*   desc :
*
================================================================*/

package logic

import (
	"container/list"
	"fmt"
)

// leetcode 20:https://leetcode.com/problems/valid-parentheses/
func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	left := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			left.PushBack(s[i])
			continue
		}
		if left.Len() != 0 && pairs[s[i]] == left.Back().Value.(byte) {
			left.Remove(left.Back())
			continue
		}
		return false
	}

	return left.Len() == 0
}

// leetcode 22: https://leetcode.com/problems/generate-parentheses/
func GenerateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	memo := make(map[int][]string)
	memo[1] = []string{"()"}

	var parenthesis func(n int) []string
	parenthesis = func(n int) []string {
		if list, exist := memo[n]; exist {
			return list
		}

		// f(n - 1) + f(1)
		results := make(map[string]struct{})
		for _, s := range parenthesis(n - 1) {
			results["("+s+")"] = struct{}{}
			results["()"+s] = struct{}{}
			results[s+"()"] = struct{}{}
		}

		// f(n-x) + f(x)
		left, right := 2, n-2
		for left <= right {
			list1 := parenthesis(left)
			list2 := parenthesis(right)
			for _, item1 := range list1 {
				for _, item2 := range list2 {
					results[item1+item2] = struct{}{}
					results[item2+item1] = struct{}{}
				}
			}
			left = left + 1
			right = right - 1
		}

		list := make([]string, 0)
		for s := range results {
			list = append(list, s)
		}

		memo[n] = list
		return list
	}

	return parenthesis(n)
}

// leetcode 29: https://leetcode.com/problems/divide-two-integers/
func Divide(dividend int, divisor int) int {

	positive := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if divisor == 0 || dividend == 0 || dividend < divisor {
		return 0
	}

	var bitcnt int
	for (divisor << (bitcnt + 1)) <= dividend {
		bitcnt = bitcnt + 1
	}

	cnt := 0
	left := dividend - (divisor << bitcnt)
	for left-divisor >= 0 {
		cnt = cnt + 1
		left = left - divisor
	}

	res := (1 << bitcnt) + cnt
	if !positive {
		res = -res
	}

	if max := (1<<31 - 1); res > max {
		return max
	}
	if min := -(1 << 31); res < min {
		return min
	}
	return res
}

// leetcode 33: https://leetcode.com/problems/search-in-rotated-sorted-array/
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
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

	return -1
}

// leetcode 329: https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])

	list := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			list = append(list, matrix[i][j])
		}
	}

	var findLongest func(matrix [][]int, from int) int

	memo := make(map[int]int)
	findLongest = func(matrix [][]int, index int) int {
		if dist, exist := memo[index]; exist {
			return dist
		}

		// 上 下
		indexs := []int{index - n, index + n}

		row := index % n
		if row-1 >= 0 { // 左
			indexs = append(indexs, index-1)
		}
		if row+1 < n { // 右
			indexs = append(indexs, index+1)
		}

		var max int = 1
		for _, nextIndex := range indexs {
			if nextIndex >= 0 && nextIndex < len(list) &&
				list[nextIndex] > list[index] {
				if dist := findLongest(matrix, nextIndex) + 1; dist > max {
					max = dist
				}
			}
		}

		memo[index] = max
		return max
	}

	max, index := 0, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index = index + 1

			if dist := findLongest(matrix, index); dist > max {
				max = dist
			}
		}
	}
	return max
}

// GCD: Greatest common divisor(最大公约数)
func MaxGCD(m, n int) int {
	mod := m % n
	if mod == 0 {
		return n
	}
	return MaxGCD(n, mod)
}

// leetcode 365: https://leetcode.com/problems/water-and-jug-problem/
// 本质上是一个加减法问题 需要注意判断无效(死循环-doingMemo、永远添加不到-最大公约数)
// 优化: getMaxGCD缓存、失败记录缓存
// 优化2: 去除最大公约数的判断反而更快一些，囧从429ms -> 338ms
func CanMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	/*
		gcdmemo := make(map[string]int)
		getMaxGCD := func(a int, b int) int {
			if a < b {
				a, b = b, a
			}
			key := strconv.Itoa(a) + "_" + strconv.Itoa(b)
			if gcd, exist := gcdmemo[key]; exist {
				return gcd
			}
			gcd := MaxGCD(a, b)
			gcdmemo[key] = gcd
			return gcd
		}
	*/

	var measure func(jug1 int, jug2 int, target int) bool

	succMemo := make(map[int]bool) // false: failed
	doingMemo := make(map[int]bool)
	measure = func(jug1 int, jug2 int, target int) bool {
		if target < 0 || jug1+jug2 < target {
			return false
		}
		if succ, exist := succMemo[target]; exist {
			return succ
		}
		if target == 0 || jug1 == target || jug2 == target || jug1+jug2 == target ||
			jug1-jug2 == target || jug2-jug1 == target {
			succMemo[target] = true
			return true
		}

		/*
			gcd := getMaxGCD(jug1, jug2)
			if gcd == 1 && targetCapacity == 1 {
				succMemo[target] = true
				return true
			}
			if gcd > 1 && target%gcd != 0 { // 这种情况倒来倒去始终为最大公约数的倍数
				succMemo[target] = false
				return false
			}
		*/

		left1, left2 := jug1-target, jug2-target
		if left1 < 0 {
			left1 = -left1
		}
		if left2 < 0 {
			left2 = -left2
		}

		result := false
		if _, exist := succMemo[left1]; !exist { // 过滤已经失败的，性能优化
			if _, exist := doingMemo[left1]; !exist { // 过滤掉正在做的，规避循环调用
				doingMemo[left1] = true
				result = measure(jug1, jug2, left1)
			}
		}
		if _, exist := succMemo[left2]; !exist {
			if _, exist := doingMemo[left2]; !exist {
				doingMemo[left2] = true
				result = result || measure(jug1, jug2, left2)
			}
		}

		succMemo[target] = result
		return result
	}

	return measure(jug1Capacity, jug2Capacity, targetCapacity)
}

// TODO
func FindKthOfTwoSortedArray(nums1 []int, nums2 []int, k int) int {
	if k <= 0 || k > len(nums1)+len(nums2) {
		return -1
	}

	l1, r1 := 0, len(nums1)-1
	l2, r2 := 0, len(nums2)-1

	var actK int = k
	for {
		fmt.Printf("nums1(%v-%v):%v nums2(%v-%v):%v k:%v\n",
			l1, r1, nums1, l2, r2, nums2, actK)

		if r1-l1 == -1 {
			return nums2[l2+actK-1]
		}
		if r2-l2 == -1 {
			return nums1[l1+actK-1]
		}
		if actK == 1 {
			return intMin(nums1[l1], nums2[l2])
		}

		if k <= 0 {
			break
		}

		var mid1, mid2 int
		if actK/2 > r1-l1+1 {
			mid1 = r1
		} else {
			mid1 = l1 + k/2 - 1
		}
		if actK/2 > r2-l2+1 {
			mid2 = r2
		} else {
			mid2 = l2 + k/2 - 1
		}
		fmt.Printf("nums1(%v-%v-%v):%v nums2(%v-%v-%v):%v k:%v\n",
			l1, mid1, r1, nums1, l2, mid2, r2, nums2, actK)

		if nums1[mid1] > nums2[mid2] {
			actK = actK - (mid2 - l2 + 1)
			l2 = mid2 + 1
		} else {
			actK = actK - (mid1 - l1 + 1)
			l1 = mid1 + 1
		}
	}

	return -1
}

func intMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 10=3 + 3 + 4 => ( 3 * 3 * 4 ) = 36
func GetMaxSplitN(n int) int {

	var fn func(n int) int

	memo := make(map[int]int)
	fn = func(n int) int {
		if n <= 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return n
		}
		if maxSplit, exist := memo[n]; exist {
			return maxSplit
		}

		var max int = n
		for i := 1; i < n; i++ {
			if maxSplit := fn(i) * fn(n-i); maxSplit > max {
				max = maxSplit
			}
		}
		memo[n] = max
		return max

	}

	return fn(n)
}

// ababc => abab c
func MaxSplitStringN(s string) []string {
	if len(s) == 0 {
		return nil
	}

	chIndexList := make(map[byte][]int) // first-end
	for i := 0; i < len(s); i++ {
		chIndexList[s[i]] = append(chIndexList[s[i]], i)
	}

	splitStrs := make([]string, 0)

	var index, left, last int = 0, 0, -1
	for index = 0; index < len(s); index++ {
		if index == last {
			splitStrs = append(splitStrs, s[left:last+1])

			left = last + 1
			continue
		}

		if lastIndexs, exist := chIndexList[s[index]]; exist && len(lastIndexs) > 1 {
			lastIndex := lastIndexs[len(lastIndexs)-1]
			if lastIndex > last {
				last = lastIndex
			}
			continue
		}

		if index > last {
			if last == 0 || last < left {
				last = left
			}
			splitStrs = append(splitStrs, s[left:last+1])

			left = last + 1
		}

	}

	return splitStrs
}

// leetcode 27: https://leetcode.com/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count = count + 1
		}
	}
	return count
}

// leetcode 338: https://leetcode.com/problems/counting-bits/submissions/
func CountBits(n int) []int {
	if n < 0 {
		return nil
	}

	// n * log(n)
	list := make([]int, n+1)
	for i := 0; i <= n; i++ {
		sum, num := 0, i
		for num != 0 {
			sum = sum + num%2
			num = num / 2
		}
		list[i] = sum
	}

	return list
}

/*
题目描述
有N个文件，每个文件的编号从0至N-1，相应大小分别记为S(i)。给定磁盘空间为C，
试实现一个函数从N个文件中选出若干个连续的文件拷贝到磁盘中，使得磁盘剩余空间最小。
示例：C, S(i) 分别为7，[1, 2, 3, 5, 4]，可求得「起始编号」、「结束编号」为 0, 2。
*/
func FindMaxSeq(c int, sizes []int) []int {
	var left, right, maxNum int

	for i := 0; i < len(sizes); i++ {
		var sum int
		for j := 0; i+j < len(sizes); j++ {
			if j == 0 {
				sum = sizes[i]
			} else {
				sum = sum + sizes[i+j]
			}

			if sum > c {
				break
			}

			if sum > maxNum {
				maxNum = sum
				left, right = i, i+j
			}
		}
	}

	return []int{left, right}
}

// leetcode 31: https://leetcode.com/problems/next-permutation/
// [1,2,3] => [1,3,2] => [3,1,2] => [2,3,1] => [3,1,2] -> [3,2,1]
func NextPermutation(nums []int) {
	var k int
	for k = len(nums) - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	if k < 0 {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
		return
	}

	var l int
	for l = len(nums) - 1; l > k; l-- {
		if nums[k] < nums[l] {
			break
		}
	}
	nums[k], nums[l] = nums[l], nums[k]

	// [k + 1:]
	start := k + 1
	for i := 0; i < (len(nums)-start)/2; i++ {
		left, right := start+i, len(nums)-1-i
		nums[left], nums[right] = nums[right], nums[left]
	}

}
