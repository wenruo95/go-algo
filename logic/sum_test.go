/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sum_test.go
*   coder: zemanzeng
*   date : 2022-02-02 21:01:54
*   desc : sum test case
*
================================================================*/

package logic

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestTreeSumClosest(t *testing.T) {

	type closestResult struct {
		nums   []int
		target int
		output int
	}

	results := []*closestResult{
		{

			nums:   []int{-1, 2, 1, -4},
			target: 1,
			output: 2,
		},
		{

			nums:   []int{0, 0, 0},
			target: 1,
			output: 0,
		},
	}

	for _, result := range results {
		output := ThreeSumClosest(result.nums, result.target)
		if output != result.output {
			t.Errorf("three_sum_closest result:%+v output:%v", result, output)
		}
	}

}

func fourSumTest(t *testing.T, fn func([]int, int) [][]int) {

	type testData struct {
		nums   []int
		target int
		arrays [][]int
	}

	datas := []*testData{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			arrays: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			arrays: [][]int{
				{2, 2, 2, 2},
			},
		},
	}

	for _, data := range datas {
		arrays := fn(data.nums, data.target)
		if !arraysEqual(arrays, data.arrays) {
			t.Errorf("four_data data:%+v arrays:%+v", data, arrays)
		}
	}

}

func arraysEqual(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[string]struct{})
	mb := make(map[string]struct{})

	for _, array := range a {
		sort.Ints(array)

		var key strings.Builder
		for index, value := range array {
			if index == 0 {
				key.WriteString(strconv.Itoa(value))
			} else {
				key.WriteString("_" + strconv.Itoa(value))
			}
		}
		ma[key.String()] = struct{}{}
	}

	for _, array := range b {
		sort.Ints(array)

		var key strings.Builder
		for index, value := range array {
			if index == 0 {
				key.WriteString(strconv.Itoa(value))
			} else {
				key.WriteString("_" + strconv.Itoa(value))
			}
		}
		mb[key.String()] = struct{}{}

		if _, exist := ma[key.String()]; !exist {
			return false
		}
	}

	return len(ma) == len(mb)
}

func TestFourSum(t *testing.T) {
	fourSumTest(t, FourSum)
}

func TestFourSum2(t *testing.T) {
	fourSumTest(t, FourSum2)
}
