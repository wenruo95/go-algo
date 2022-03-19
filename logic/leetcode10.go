/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic01.go
*   coder: zemanzeng
*   date : 2022-03-19 22:06:18
*   desc : leetcode 10~19
*
================================================================*/

package logic

import (
	"sort"
	"strconv"
	"strings"
)

// leetcode 10: https://leetcode.com/problems/regular-expression-matching/
func RegularIsMatch(s string, p string) bool {
	memo := make(map[string]bool)

	var regularIsMatch func(string, int, string, int) bool
	regularIsMatch = func(s string, index int, p string, pindex int) bool {
		if len(p) == pindex {
			return len(s) == index
		}
		if len(s) == index {
			if (len(p)-pindex)%2 == 1 {
				return false
			}
			for i := pindex; i+1 < len(p); i = i + 2 {
				if p[i+1] != '*' {
					return false
				}
			}
			return true
		}

		key := strconv.Itoa(index) + "_" + strconv.Itoa(pindex)
		if value, exist := memo[key]; exist {
			return value
		}

		var match bool
		if s[index] == p[pindex] || p[pindex] == '.' {
			if pindex+1 < len(p) && p[pindex+1] == '*' {
				match = regularIsMatch(s, index, p, pindex+2) || // 0次
					regularIsMatch(s, index+1, p, pindex) // 1次 或 多次
			} else {
				match = regularIsMatch(s, index+1, p, pindex+1)
			}
		} else {
			if pindex+1 < len(p) && p[pindex+1] == '*' {
				match = regularIsMatch(s, index, p, pindex+2)
			} else {
				match = false
			}
		}

		memo[key] = match
		return match
	}

	return regularIsMatch(s, 0, p, 0)
}

// leetcode 16: https://leetcode.com/problems/3sum-closest/
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var mindist int
	var indexs [3]int
	for i := 0; i < len(nums)-2; i++ {
		i1, i2 := twoSumClosest(nums, i+1, target-nums[i])
		if distance := abs(nums[i]+nums[i1]+nums[i2], target); i == 0 || distance < mindist {
			mindist = distance
			indexs[0], indexs[1], indexs[2] = i, i1, i2
		}
	}

	return nums[indexs[0]] + nums[indexs[1]] + nums[indexs[2]]
}

// leetcode 17: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func LetterCombinations(digits string) []string {
	digitStr := [10][]string{
		0: {},
		1: {},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	var indexs []int
	for _, digit := range digits {
		indexs = append(indexs, int(digit-'0'))
	}

	var fn func(indexs []int, cur int) []string
	fn = func(indexs []int, cur int) []string {
		if cur == len(indexs) {
			return nil
		}

		strs := digitStr[indexs[cur]]
		afStrs := fn(indexs, cur+1)
		if len(afStrs) == 0 {
			return strs
		}

		list := make([]string, 0)
		for _, str := range strs {
			for _, afStr := range afStrs {
				list = append(list, str+afStr)
			}
		}

		return list
	}

	return fn(indexs, 0)
}

func abs(num, target int) int {
	if num > target {
		return num - target
	}
	return target - num
}

func twoSumClosest(nums []int, start int, target int) (int, int) {
	low, high := start, len(nums)-1

	i1, i2 := low, high
	min := abs(nums[low]+nums[high], target)
	for low < high {
		sum := nums[low] + nums[high]
		if distance := abs(sum, target); distance < min {
			min = distance
			i1, i2 = low, high
		}

		if sum > target {
			high = high - 1
		} else if sum == target {
			break
		} else {
			low = low + 1
		}
	}
	return i1, i2
}

// leetcode 18: https://leetcode.com/problems/4sum/
func FourSum(nums []int, target int) [][]int {

	cnts := make(map[int]int)
	for _, num := range nums {
		cnts[num] = cnts[num] + 1
	}

	var fn func(target int, count int) [][]int

	fn = func(target int, count int) [][]int {

		if count == 2 {
			numSet := make(map[int]struct{})
			for num := range cnts {
				if cnts[num] > 0 && cnts[target-num] > 0 && (num != target-num || cnts[num] > 1) {
					if _, exist := numSet[target-num]; exist {
						continue
					}
					numSet[num] = struct{}{}
				}
			}

			result := make([][]int, 0)
			for num := range numSet {
				result = append(result, []int{num, target - num})
			}
			return result
		}

		result := make([][]int, 0)
		noRepeatSet := make(map[string]struct{})
		for num, cnt := range cnts {
			if cnt == 0 {
				continue
			}

			cnts[num] = cnt - 1
			arrays := fn(target-num, count-1)
			cnts[num] = cnt

			for _, array := range arrays {
				array = append(array, num)
				if count-1 == 2 {
					result = append(result, array)
					continue
				}

				sort.Ints(array)

				var key strings.Builder
				for index, value := range array {
					if index == 0 {
						key.WriteString(strconv.Itoa(value))
					} else {
						key.WriteString("_" + strconv.Itoa(value))
					}
				}
				if _, exist := noRepeatSet[key.String()]; exist {
					continue
				}
				noRepeatSet[key.String()] = struct{}{}

				result = append(result, array)
			}
		}

		return result
	}

	return fn(target, 4)
}

func FourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)

	var fn func(n int, start int, target int) [][]int

	fn = func(n int, start int, target int) [][]int {

		records := make([][]int, 0)

		if n < 2 {
			return records
		}

		if n > 2 {
			for i := start; i < len(nums); i++ {
				lists := fn(n-1, i+1, target-nums[i])
				for j := 0; j < len(lists); j++ {
					lists[j] = append([]int{nums[i]}, lists[j]...)
				}
				records = append(records, lists...)

				for i+1 < len(nums) && nums[i] == nums[i+1] {
					i = i + 1
				}

			}
			return records
		}

		low, high := start, len(nums)-1
		for low < high {

			num := nums[low] + nums[high]
			left, right := nums[low], nums[high]

			if num < target {
				for low < high && nums[low] == left {
					low = low + 1
				}

			} else if num > target {
				for low < high && nums[high] == right {
					high = high - 1
				}

			} else {
				records = append(records, []int{left, right})

				for low < high && nums[low] == left {
					low = low + 1
				}
				for low < high && nums[high] == right {
					high = high - 1
				}
			}

		}

		return records
	}

	return fn(4, 0, target)
}
