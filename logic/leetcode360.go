/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode360.go
*   coder: zemanzeng
*   date : 2022-03-19 22:02:32
*   desc : leetcode 360~369
*
================================================================*/

package logic

// leetcode 365: https://leetcode.com/problems/water-and-jug-problem/
func CanMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {

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
