/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sum.go
*   coder: zemanzeng
*   date : 2022-02-02 21:01:05
*   desc : sum
*
================================================================*/

package logic

import (
	"sort"
	"strconv"
	"strings"
)

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

// leetcode 18: https://leetcode.com/problems/4sum/
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

// leetcode 39: https://leetcode.com/problems/combination-sum/
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Ints(candidates)

	memo := make(map[int][][]int)

	var fn func(target int) [][]int
	fn = func(target int) [][]int {
		if target < candidates[0] {
			return nil
		}
		if v, exist := memo[target]; exist {
			return v
		}

		arrays := make([][]int, 0)
		for i := 0; i < len(candidates); i++ {
			if target == candidates[i] {
				arrays = append(arrays, []int{candidates[i]})
				continue
			}

			lists := fn(target - candidates[i])
			for _, list := range lists {
				if len(list) > 0 && candidates[i] > list[0] { // 确保数据有序
					continue
				}
				arrays = append(arrays,
					append([]int{candidates[i]}, list...),
				)
			}
		}

		memo[target] = arrays
		return arrays
	}

	return fn(target)
}

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
