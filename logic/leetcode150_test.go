/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode150_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:31:58
*   desc :
*
================================================================*/

package logic

import "testing"

func TestReverseWords(t *testing.T) {
	type reverseResult struct {
		input  string
		output string
	}

	results := []*reverseResult{
		{
			input:  "the sky is blue",
			output: "blue is sky the",
		},
		{
			input:  "  hello world  ",
			output: "world hello",
		},
		{
			input:  "a good   example",
			output: "example good a",
		},
	}

	for _, result := range results {
		if output := ReverseWords(result.input); output != result.output {
			t.Errorf("reverse_words result:%+v output[%v]:%v expect[%v]:%v",
				result, len(output), output, len(result.output), result.output)
		}
	}

}

func TestMaxProduct(t *testing.T) {
	type testData struct {
		nums   []int
		target int
	}

	datas := []*testData{
		{nums: []int{-1, 0, -2}, target: 0},
		{nums: []int{2, 3, -2, 4}, target: 6},
		{nums: []int{-1, 2, 3, 2, 4, -2}, target: 96},
	}
	for _, data := range datas {
		if target := MaxProduct(data.nums); data.target != target {
			t.Errorf("max_product nums:%+v target:%v expect:%v", data.nums, target, data.target)
		}
	}
}

func TestFindMinInRotatedArray(t *testing.T) {
	type testData struct {
		nums   []int
		target int
	}

	datas := []*testData{
		{nums: []int{1, 2, 3, 4, 5}, target: 1},
		{nums: []int{5, 1, 2, 3, 4}, target: 1},
		{nums: []int{4, 5, 1, 2, 3}, target: 1},
		{nums: []int{3, 4, 5, 1, 2}, target: 1},
		{nums: []int{2, 3, 4, 5, 1}, target: 1},

		{nums: []int{1, 2, 3, 4, 5, 6}, target: 1},
		{nums: []int{6, 1, 2, 3, 4, 5}, target: 1},
		{nums: []int{5, 6, 1, 2, 3, 4}, target: 1},
		{nums: []int{4, 5, 6, 1, 2, 3}, target: 1},
		{nums: []int{3, 4, 5, 6, 1, 2}, target: 1},
		{nums: []int{2, 3, 4, 5, 6, 1}, target: 1},
	}
	for _, data := range datas {
		if target := FindMinItemInRotatedArray(data.nums); target != data.target {
			t.Errorf("find_min_in_rorated_array nums:%+v target:%v expect:%v", data.nums, target, data.target)
		}
	}
}
