/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sort_test.go
*   coder: zemanzeng
*   date : 2022-02-12 11:40:22
*   desc :
*
================================================================*/

package logic

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type sortResult struct {
	nums []int
}

func newSortResults(autoCase bool) []*sortResult {
	// 人工测试用例
	results := []*sortResult{
		{nums: []int{}},
		{nums: []int{0}},
		{nums: []int{1, 0}},
		{nums: []int{1, 1, 1}},
		{nums: []int{9, 2, 1, 4, 7, 5, 3, -1, 4}},
		{nums: []int{9, 4, 6, 8, 3, 10, 4, 6}},
	}
	if !autoCase {
		return results
	}

	// 自动生成的测试用例
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 1000; i++ {
		nums := make([]int, i)
		for j := 0; j < i; j++ {
			nums[j] = rand.Int()
		}
		results = append(results, &sortResult{nums: nums})
	}
	return results
}

func intListEqual(dst, src []int) bool {
	if len(dst) != len(src) {
		return false
	}
	for i := 0; i < len(dst); i++ {
		if dst[i] != src[i] {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	results := newSortResults(true)
	for _, result := range results {
		expectInts := make([]int, len(result.nums))
		copy(expectInts, result.nums)
		sort.Ints(expectInts)

		QuickSort(result.nums)
		if !intListEqual(result.nums, expectInts) {
			t.Errorf("quick_sort nums:%+v expect:%+v", result.nums, expectInts)
		}
	}

}

func TestFindKth(t *testing.T) {
	type testData struct {
		nums   []int
		k      int
		result int
	}

	datas := []*testData{
		{
			nums:   []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
			k:      3,
			result: 9,
		},
		{
			nums:   []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
			k:      10,
			result: 4,
		},
		{
			nums:   []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
			k:      12,
			result: 2,
		},
	}

	for _, data := range datas {
		if result := FindKth(data.nums, data.k); result != data.result {
			t.Errorf("find_k_th data:%+v result:%v", data, result)
		}
	}
}
