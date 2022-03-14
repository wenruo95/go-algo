/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic_test.go
*   coder: zemanzeng
*   date : 2022-02-03 01:16:43
*   desc :
*
================================================================*/

package logic

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	chs := "abcedfghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	uniqs := "=!@#$"
	collects := chs + strings.ToUpper(chs) + numbers + uniqs

	var s string
	for i := 0; i < 12; i++ {
		s = s + string(collects[rand.Int()%len(collects)])
	}
	t.Logf("rand:%v", s)

}

func TestIsValidParentheses(t *testing.T) {
	type parenthesesResult struct {
		s     string
		valid bool
	}
	results := []*parenthesesResult{
		{
			s:     "()",
			valid: true,
		},
		{
			s:     "()[]{}",
			valid: true,
		},
		{
			s:     "(]",
			valid: false,
		},
	}

	for _, result := range results {
		if valid := IsValidParentheses(result.s); valid != result.valid {
			t.Errorf("is_valid_parentheses result:%+v valid:%v", result, valid)
		}
	}

}

func TestGenerateParenthesis(t *testing.T) {
	type generateParenthesis struct {
		n      int
		output []string
	}

	results := []*generateParenthesis{
		{
			n:      1,
			output: []string{"()"},
		},
		{
			n:      2,
			output: []string{"()()", "(())"},
		},
		{
			n:      3,
			output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n: 4,
			output: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())",
				"(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	for _, result := range results {
		output := GenerateParenthesis(result.n)
		if !stringListItemEqual(output, result.output) {
			t.Errorf("generate_parenthesis result:%+v output:%+v", result, output)
		}

	}

}

func TestDivide(t *testing.T) {
	type divideResult struct {
		dividend int
		divisor  int
		quotient int
	}

	results := []*divideResult{
		{
			dividend: 10,
			divisor:  3,
			quotient: 3,
		},
		{
			dividend: 7,
			divisor:  -3,
			quotient: -2,
		},
		{
			dividend: 1,
			divisor:  2,
			quotient: 0,
		},
		{
			dividend: -2147483648,
			divisor:  -1,
			quotient: 2147483647,
		},
	}

	for _, result := range results {
		quotient := Divide(result.dividend, result.divisor)
		if quotient != result.quotient {
			t.Errorf("divide result:%+v quotient:%v", result, quotient)
		}
	}

}

func TestSearch(t *testing.T) {
	type searchResult struct {
		nums   []int
		target int
		index  int
	}

	results := []*searchResult{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			index:  4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			index:  -1,
		},
		{
			nums:   []int{1},
			target: 0,
			index:  -1,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 4,
			index:  3,
		},
		{
			nums:   []int{3, 1},
			target: 1,
			index:  1,
		},
	}
	for _, result := range results {
		index := SearchInRotatedSortedArray(result.nums, result.target)
		if index != result.index {
			t.Errorf("search_in_rotated_sorted_array result:%+v index:%v", result, index)
		}
	}

}

func TestLongestIncreasingPath(t *testing.T) {
	type longestPathResult struct {
		matrix [][]int
		output int
	}

	results := []*longestPathResult{
		{
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			output: 4,
		},
		{
			matrix: [][]int{
				{6, 8},
				{7, 2},
			},
			output: 2,
		},
	}
	for _, result := range results {
		if output := LongestIncreasingPath(result.matrix); output != result.output {
			t.Errorf("longest_increasing_path result:%+v output:%v", result, output)
		}
	}

}

func TestCanMeasureWater(t *testing.T) {
	type measureResult struct {
		jug1   int
		jug2   int
		target int
		output bool
	}

	results := []*measureResult{
		{jug1: 3, jug2: 5, target: 4, output: true},
		{jug1: 2, jug2: 6, target: 5, output: false},
		{jug1: 1, jug2: 2, target: 3, output: true},
		{jug1: 9, jug2: 6, target: 1, output: false},
		{jug1: 10000, jug2: 10001, target: 1, output: true},
		{jug1: 4, jug2: 6, target: 8, output: true},
		{jug1: 34, jug2: 5, target: 6, output: true},
		{jug1: 1, jug2: 1, target: 12, output: false},
	}

	for _, result := range results {
		output := CanMeasureWater(result.jug1, result.jug2, result.target)
		if output != result.output {
			t.Errorf("can_measure_water result:%+v output:%v", result, output)
		}
	}

}

// TODO
func TestKthOfTwoSortedArray(t *testing.T) {

	type testData struct {
		nums1 []int
		nums2 []int
		k     int
	}

	datas := []*testData{
		{
			nums1: []int{1, 2, 5, 8, 10},
			nums2: []int{3, 6, 7, 9, 10},
			k:     5,
		},
		{
			nums1: []int{1, 2, 5, 5, 8, 10},
			nums2: []int{3, 6, 7, 9, 10},
			k:     5,
		},
		{
			nums1: []int{1, 2, 5, 5, 8, 10},
			nums2: []int{3, 6, 7, 9, 10},
			k:     11,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{2},
			k:     1,
		},
	}
	if len(datas) < 100 {
		return
	}

	for _, data := range datas {
		arr := make([]int, 0)
		arr = append(arr, data.nums1...)
		arr = append(arr, data.nums2...)
		sort.Ints(arr)
		if data.k < 0 || data.k > len(arr) {
			t.Errorf("invalid k:%v len:%v", data.k, len(arr))
		}

		result1 := FindKthOfTwoSortedArray(data.nums1, data.nums2, data.k)
		result2 := FindKthOfTwoSortedArray(data.nums2, data.nums1, data.k)

		if result1 != result2 || result1 != arr[data.k-1] {
			t.Errorf("data:%+v result1:%v result2:%v expect:%v", data, result1, result2, arr[data.k-1])

		}

	}

}

func TestMaxSplitN(t *testing.T) {

	type testData struct {
		n   int
		max int
	}

	datas := []*testData{
		{
			n:   0,
			max: 0,
		},
		{
			n:   10,
			max: 36,
		},
		{
			n:   20,
			max: 1458,
		},
	}

	for _, data := range datas {
		if max := GetMaxSplitN(data.n); max != data.max {
			t.Errorf("get_max_split_n data:%+v max:%v", data, max)
		}
	}

}

func TestMaxSplitStringN(t *testing.T) {

	type testData struct {
		s      string
		output []string
	}

	datas := []*testData{
		{
			s:      "ababc",
			output: []string{"abab", "c"},
		},
	}

	for _, data := range datas {
		output := MaxSplitStringN(data.s)
		if !stringListEqual(output, data.output) {
			t.Errorf("max_split_string_n data:%+v output:%v", data, output)
		}
	}

}

func TestRemoveElement(t *testing.T) {
	type testData struct {
		nums  []int
		val   int
		count int
	}

	datas := []*testData{
		{
			nums:  []int{3, 2, 2, 3},
			val:   3,
			count: 2,
		},
		{
			nums:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:   2,
			count: 5,
		},
	}

	for _, data := range datas {
		if count := RemoveElement(data.nums, data.val); count != data.count {
			t.Errorf("remove_element data:%+v count:%v", data, count)
		}
	}

}

func TestCountBits(t *testing.T) {
	type testData struct {
		n    int
		list []int
	}

	datas := []*testData{
		{
			n:    2,
			list: []int{0, 1, 1},
		},
		{
			n:    5,
			list: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, data := range datas {
		if list := CountBits(data.n); !intListEqual(list, data.list) {
			t.Errorf("count bits error. data:%+v list:%v", data, list)
		}
	}

}

func TestFindMaxSeq(t *testing.T) {
	type testData struct {
		c      int
		sizes  []int
		result []int
	}

	datas := []*testData{
		{
			c:      1,
			sizes:  []int{1, 2, 3, 5, 4},
			result: []int{0, 0},
		},
		{
			c:      7,
			sizes:  []int{1, 2, 3, 5, 4},
			result: []int{0, 2},
		},
		{
			c:      7,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{5, 5},
		},
		{
			c:      21,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{1, 5},
		},
		{
			c:      22,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{0, 5},
		},
		{
			c:      23,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{0, 5},
		},
	}

	for _, data := range datas {
		result := FindMaxSeq(data.c, data.sizes)
		if len(result) != 2 ||
			result[0] != data.result[0] || result[1] != data.result[1] {
			t.Errorf("find_max_seq error. data:%+v result:%v", data, result)
		}

	}
}

func TestNextPermutation(t *testing.T) {
	type testData struct {
		nums   []int
		output []int
	}

	datas := []*testData{
		{
			nums:   []int{1, 2, 3},
			output: []int{1, 3, 2},
		},
		{
			nums:   []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		{
			nums:   []int{1, 3, 2},
			output: []int{2, 1, 3},
		},
		{
			nums:   []int{1, 1, 5},
			output: []int{1, 5, 1},
		},
		{
			nums:   []int{5, 4, 7, 5, 3, 2},
			output: []int{5, 5, 2, 3, 4, 7},
		},
	}

	for _, data := range datas {
		nums := make([]int, len(data.nums))
		for index, num := range data.nums {
			nums[index] = num
		}

		NextPermutation(nums)
		if !intListEqual(nums, data.output) {
			t.Errorf("next_permutation error. data:%+v nums:%+v", data, nums)
		}
	}
}

func TestLongestValidParentheses(t *testing.T) {
	type testData struct {
		s   string
		num int
	}

	datas := []*testData{
		{
			s:   "(()",
			num: 2,
		},
		{
			s:   ")()())",
			num: 4,
		},
		{
			s:   "",
			num: 0,
		},
	}

	for _, data := range datas {
		if num := LongestValidParentheses(data.s); num != data.num {
			t.Errorf("longest_valid_parentheses error. data:%+v num:%v", data, num)
		}
	}

}
