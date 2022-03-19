/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode30_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:25:58
*   desc :
*
================================================================*/

package logic

import (
	"strings"
	"testing"
)

func testFindSubString(t *testing.T, fn func(s string, words []string) []int) {
	type testData struct {
		s      string
		words  []string
		indexs []int
	}

	datas := []*testData{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			indexs: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			indexs: []int{},
		},
		{
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			indexs: []int{6, 9, 12},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "good"},
			indexs: []int{8},
		},
	}

	{
		s := strings.Repeat("a", 5000+10)
		words := make([]string, 5000)
		for i := 0; i < 5000; i++ {
			words[i] = "a"
		}
		datas = append(datas, &testData{s: s, words: words, indexs: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
	}

	for _, data := range datas {
		if indexs := fn(data.s, data.words); !intListEqual(indexs, data.indexs) {
			t.Errorf("find_sub_string error. data:%+v indexs:%v", data, indexs)
		}
	}
}

// leetcode 30
func TestFindSubString(t *testing.T) {
	testFindSubString(t, FindSubstring)
}

func TestFindSubString2(t *testing.T) {
	testFindSubString(t, FindSubstring2)
}

// leetcode 31
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

// leetcode 33
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

// leetcode 33
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

// leetcode 34
func TestSearchRange(t *testing.T) {
	type testData struct {
		nums   []int
		target int
		result []int
	}

	datas := []*testData{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			result: []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			result: []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			result: []int{-1, -1},
		},
	}

	for _, data := range datas {
		if result := SearchRange(data.nums, data.target); !intListEqual(result, data.result) {
			t.Errorf("search_range error. data:%+v result:%v", data, result)
		}
	}

}

// leetcode 35
func TestSearchInsert(t *testing.T) {
	type testData struct {
		nums   []int
		target int
		output int
	}

	datas := []*testData{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
	}

	for _, data := range datas {
		if output := SearchInsert(data.nums, data.target); output != data.output {
			t.Errorf("search_insert error. data:%+v output:%v", data, output)
		}
	}

}

// leetcode 36
func TestIsValidSudoku(t *testing.T) {

	type testData struct {
		board [][]byte
		valid bool
	}

	datas := []*testData{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			valid: true,
		},
		{
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			valid: false,
		},
	}

	for _, data := range datas {
		if valid := IsValidSudoku(data.board); valid != data.valid {
			t.Errorf("is_valid_sudoku error. valid:%v expect:%v", valid, data.valid)
		}
	}
}

// leetcode 37
func TestSolveSudoku(t *testing.T) {
	type testData struct {
		board [][]byte
	}

	datas := []*testData{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
	}

	for _, data := range datas {
		SolveSudoku(data.board)
		if !IsValidSudoku(data.board) {
			t.Errorf("solve_sudo_ku error")
		}
		//t.Logf("board:%+v", board2str(data.board))
	}

}

func board2str(board [][]byte) string {
	var s string = "\n"
	for _, bts := range board {
		for _, bt := range bts {
			s = s + string(bt)
		}
		s = s + "\n"
	}
	return s
}

// leetcode 38
func TestCountAndSay(t *testing.T) {
	type testData struct {
		n int
		s string
	}

	datas := []*testData{
		{
			n: 1,
			s: "1",
		},
		{
			n: 4,
			s: "1211",
		},
	}

	for _, data := range datas {
		if s := CountAndSay(data.n); s != data.s {
			t.Errorf("count_and_say error. data:%+v s:%v", data, s)
		}
	}

}
