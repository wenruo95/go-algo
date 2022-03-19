/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode10_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:22:14
*   desc :
*
================================================================*/

package logic

import "testing"

// leetcode 10
func TestRegularIsMatch(t *testing.T) {
	type regularResult struct {
		s     string
		p     string
		match bool
	}

	results := []*regularResult{
		{
			s:     "aa",
			p:     "a",
			match: false,
			// Explanation: "a" does not match the entire string "aa".
		},
		{
			s:     "aa",
			p:     "a*",
			match: true,
			// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
		},
		{
			s:     "ab",
			p:     ".*",
			match: true,
			// Explanation: ".*" means "zero or more (*) of any character (.)".
		},
		{
			s:     "aab",
			p:     "c*a*b",
			match: true,
		},
		{
			s:     "mississippi",
			p:     "mis*is*p*.",
			match: false,
		},
		{
			s:     "aaa",
			p:     "ab*a*c*a",
			match: true,
		},
	}

	for _, result := range results {
		if match := RegularIsMatch(result.s, result.p); match != result.match {
			t.Errorf("regular_is_match result:%+v match:%v", result, match)
		}
	}
}

// leetcode 16
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

// leetcode 17
func TestLetterCombinations(t *testing.T) {

	type letterResult struct {
		digits string
		output []string
	}

	results := []*letterResult{
		{
			digits: "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			output: []string{},
		},
		{
			digits: "2",
			output: []string{"a", "b", "c"},
		},
	}

	for _, result := range results {
		combs := LetterCombinations(result.digits)
		if !stringListItemEqual(combs, result.output) {
			t.Errorf("letter_combineations. digits:%+v output:%+v combs:%+v",
				result.digits, result.output, combs)
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

// leetcode 18
func TestFourSum(t *testing.T) {
	fourSumTest(t, FourSum)
}

func TestFourSum2(t *testing.T) {
	fourSumTest(t, FourSum2)
}
