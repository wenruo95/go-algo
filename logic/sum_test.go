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

func TestFourSum(t *testing.T) {
	fourSumTest(t, FourSum)
}

func TestFourSum2(t *testing.T) {
	fourSumTest(t, FourSum2)
}

func TestCombinationSum(t *testing.T) {
	type testData struct {
		candidates []int
		target     int
		arrays     [][]int
	}

	datas := []*testData{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			arrays: [][]int{
				{2, 2, 3}, {7},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			arrays: [][]int{
				{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
			},
		},
		{
			candidates: []int{2},
			target:     1,
			arrays:     [][]int{},
		},
		{
			candidates: []int{3, 12, 9, 11, 6, 7, 8, 5, 4},
			target:     15,
			arrays: [][]int{
				{3, 3, 3, 3, 3},
				{3, 3, 3, 6},
				{3, 3, 4, 5},
				{3, 3, 9},
				{3, 4, 4, 4},
				{3, 4, 8},
				{3, 5, 7},
				{3, 6, 6},
				{3, 12},
				{4, 4, 7},
				{4, 5, 6},
				{4, 11},
				{5, 5, 5},
				{6, 9},
				{7, 8},
			},
		},
	}

	for _, data := range datas {
		if arrays := CombinationSum(data.candidates, data.target); !arraysEqual(arrays, data.arrays) {
			t.Errorf("combination_sum error. data:%+v arrays:%v", data, arrays)
		}
	}

}
