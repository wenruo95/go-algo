package logic

import (
	"container/heap"
	"sort"
)

// leetcode 347: https://leetcode.com/problems/top-k-frequent-elements/
func TopKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)

	h := make(FrequentHeap, 0)
	heap.Init(&h)

	value, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if value == nums[i] {
			count = count + 1
			continue
		}
		heap.Push(&h, &FrequentData{Count: count, Value: value})

		value = nums[i]
		count = 1
	}
	if len(nums) > 0 {
		heap.Push(&h, &FrequentData{Count: count, Value: value})
	}

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(*FrequentData).Value)
	}
	return res
}
