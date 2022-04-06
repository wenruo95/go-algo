/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode40.go
*   coder: zemanzeng
*   date : 2022-03-19 22:00:37
*   desc : leetcode 40~49
*
================================================================*/

package logic

import (
	"sort"
	"strconv"
	"strings"
)

// leetcode 40: https://leetcode.com/problems/combination-sum-ii/
func CombinationSum2(candidates []int, target int) [][]int {

	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)

	memo := make(map[string][][]int)

	var fn func(index int, target int) [][]int
	fn = func(index, target int) [][]int {
		if index >= len(candidates) || target < candidates[index] {
			return nil
		}
		key := strconv.Itoa(index) + "_" + strconv.Itoa(target)
		if v, exist := memo[key]; exist {
			return v
		}

		arrays := make([][]int, 0)
		for i := index; i < len(candidates); i++ {
			if i-1 >= index && candidates[i-1] == candidates[i] {
				continue
			}

			if target == candidates[i] {
				arrays = append(arrays, []int{candidates[i]})
				continue
			}

			lists := fn(i+1, target-candidates[i])
			for _, list := range lists {
				if len(list) > 0 && candidates[i] > list[0] { // 确保数据有序
					continue
				}
				arrays = append(arrays,
					append([]int{candidates[i]}, list...),
				)
			}
		}

		memo[key] = arrays
		return arrays
	}

	return fn(0, target)
}

// leetcode 41: https://leetcode.com/problems/first-missing-positive/
func FirstMissingPositive(nums []int) int {
	sort.Ints(nums)

	var firstPositive int = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && firstPositive == -1 {
			firstPositive = nums[i]
			if firstPositive != 1 {
				break
			}
		}

		if i-1 >= 0 && nums[i]-nums[i-1] > 1 && nums[i-1] > 0 {
			return nums[i-1] + 1
		}
	}
	if firstPositive != 1 {
		return 1
	}

	return nums[len(nums)-1] + 1
}

func FirstMissingPositive2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i]-1] != nums[i] {
			tmp := nums[i] - 1
			nums[i], nums[tmp] = nums[tmp], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

// leetcode 42: https://leetcode.com/problems/trapping-rain-water/
func Trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	var max int = -1
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		left[i] = max
	}

	max = -1
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		right[i] = max
	}

	var sum int
	for i := 0; i < len(height); i++ {
		if left[i] > height[i] && right[i] > height[i] {
			sum = sum + intMin(left[i]-height[i], right[i]-height[i])
		}
	}
	return sum
}

// leetcode 43: https://leetcode.com/problems/multiply-strings/
func Multiply(num1 string, num2 string) string {
	var sum string = "0"
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			num := int(num1[i]-'0') * int(num2[j]-'0')
			if num > 0 {
				zeroCount := (len(num1) - 1 - i) + (len(num2) - 1 - j)
				sum = StrAdd(sum, strconv.Itoa(num)+strings.Repeat("0", zeroCount))
			}
		}
	}
	return sum
}

func StrAdd(num1 string, num2 string) string {
	var bts []byte
	if len(num1) > len(num2) {
		bts = make([]byte, len(num1)+1)
	} else {
		bts = make([]byte, len(num2)+1)
	}

	i1, i2 := len(num1)-1, len(num2)-1
	pos, high := len(bts)-1, 0
	for i1 >= 0 || i2 >= 0 || high > 0 {
		var num int
		if i1 >= 0 {
			num = num + int(num1[i1]-'0')
		}
		if i2 >= 0 {
			num = num + int(num2[i2]-'0')
		}

		num = num + high
		high = num / 10
		bts[pos] = byte(num%10) + '0'

		i1 = i1 - 1
		i2 = i2 - 1
		pos = pos - 1
	}

	if pos == 0 {
		return string(bts[1:])
	}
	return string(bts)
}

func Multiply2(num1 string, num2 string) string {

	arr := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p1 := i + j + 1
			num := int(num1[i]-'0')*int(num2[j]-'0') + arr[p1]

			arr[p1] = num % 10
			arr[p1-1] = arr[p1-1] + num/10
		}
	}

	sum := &strings.Builder{}
	sum.Grow(len(arr))
	for i := 0; i < len(arr); i++ {
		if sum.Len() == 0 && arr[i] == 0 {
			continue
		}
		sum.WriteByte(byte(arr[i]) + '0')
	}

	if sum.Len() == 0 {
		return "0"
	}
	return sum.String()
}

// leetcode 44: https://leetcode.com/problems/wildcard-matching/
func IsMatch(s string, p string) bool {
	var index, pindex int
	var star, backtrackIdx int = -1, -1
	for index < len(s) {
		if pindex < len(p) && (s[index] == p[pindex] || p[pindex] == '?') {
			index = index + 1
			pindex = pindex + 1
			continue
		}

		if pindex < len(p) && p[pindex] == '*' { // 默认忽略*
			star = pindex
			backtrackIdx = index
			pindex = pindex + 1
			continue
		}

		if star >= 0 && p[star] == '*' { // 回溯
			backtrackIdx = backtrackIdx + 1
			index = backtrackIdx
			pindex = star
			continue
		}

		return false
	}

	for pindex < len(p) {
		if p[pindex] == '*' {
			pindex = pindex + 1
			continue
		}
		break
	}

	return pindex == len(p)
}

// leetcode 45: https://leetcode.com/problems/jump-game-ii/
func JumpGameII(nums []int) int {
	memo := make(map[int]int)

	var minJumpFn func(left int) int
	minJumpFn = func(left int) int {
		if left >= len(nums)-1 {
			return 0
		}
		if v, exist := memo[left]; exist {
			return v
		}

		var min int = -1
		for step := 1; step <= nums[left]; step++ {
			if m := minJumpFn(left + step); min == -1 || m < min {
				min = m
			}
		}
		if min == -1 {
			return minJumpFn(left+1) + 1
		}
		min = min + 1

		memo[left] = min
		return min
	}

	return minJumpFn(0)
}

func JumpGameII2(nums []int) int {
	var step, curFarthest, curEnd int

	for i := 0; i < len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i + nums[i]
		}

		if i == curEnd {
			step = step + 1
			curEnd = curFarthest
		}
	}

	return step
}

// leetcode 46: https://leetcode.com/problems/permutations/
func Permute(nums []int) [][]int {

	var fn func(begin int)

	arrays := make([][]int, 0)
	fn = func(begin int) {
		if begin >= len(nums) {
			arrays = append(arrays, append([]int{}, nums[0:]...))
			return
		}

		for i := begin; i < len(nums); i++ {
			nums[i], nums[begin] = nums[begin], nums[i]
			fn(begin + 1)
			nums[begin], nums[i] = nums[i], nums[begin]
		}

	}

	fn(0)
	return arrays
}

// leetcode 47: https://leetcode.com/problems/permutations-ii/
// ideas refer from labuladong: https://labuladong.gitee.io/algo/4/30/109/
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var fn func()

	arrays := make([][]int, 0)

	used := make([]bool, len(nums))
	track := make([]int, 0)
	fn = func() {
		if len(track) >= len(nums) {
			arrays = append(arrays, append([]int{}, track...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			track = append(track, nums[i])

			fn()

			track = track[0 : len(track)-1]
			used[i] = false
		}

	}

	fn()
	return arrays
}

// leetcode 48: https://leetcode.com/problems/rotate-image/
func Rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	// 第i层
	for i := 0; i < len(matrix)/2; i++ {

		// 左上角元素
		for j := i; j <= len(matrix)-i-2; j++ {

			var tmp, row, column int
			// 右下: row, column => row + (j - i), len(matrix) -1 - i
			{
				row, column = i, j
				tmp, matrix[row+j-i][len(matrix)-1-i] = matrix[row+j-i][len(matrix)-1-i], matrix[row][column]
			}

			// 下左 row, column => len(matrix) -1 - i, column - (j - i)
			{
				row, column = row+j-i, len(matrix)-1-i
				matrix[len(matrix)-1-i][column-(j-i)], tmp = tmp, matrix[len(matrix)-1-i][column-(j-i)]
			}

			// 左上 row, column => len(matrix) - ( j - i ), i
			{
				row, column = len(matrix)-1-i, column-(j-i)
				matrix[row-(j-i)][i], tmp = tmp, matrix[row-(j-i)][i]
			}

			// 上右
			{
				row, column = row-(j-i), i
				matrix[i][j] = tmp
			}

		}
	}

}

// leetcode 49: https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]int)

	for index, str := range strs {
		var array [26]int
		for _, c := range str {
			array[c-'a']++
		}
		if v, exist := groups[array]; exist {
			groups[array] = append(v, index)
		} else {
			groups[array] = []int{index}
		}
	}

	arrays := make([][]string, 0)
	for _, indexs := range groups {
		array := make([]string, len(indexs))
		for i, index := range indexs {
			array[i] = strs[index]
		}
		arrays = append(arrays, array)
	}

	return arrays
}
