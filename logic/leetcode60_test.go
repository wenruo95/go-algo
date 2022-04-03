/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode60_test.go
*   coder: zemanzeng
*   date : 2022-03-29 10:37:31
*   desc : leetcode 60~69 test case
*
================================================================*/

package logic

import (
	"strings"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	type testData struct {
		n           int
		k           int
		permutation string
	}

	datas := []*testData{
		{
			n:           3,
			k:           3,
			permutation: "213",
		},
		{
			n:           4,
			k:           9,
			permutation: "2314",
		},
		{
			n:           3,
			k:           1,
			permutation: "123",
		},
	}

	for _, data := range datas {
		if permutation := GetPermutation(data.n, data.k); permutation != data.permutation {
			t.Errorf("get_permutation error. data:%+v permutation:%v", data, permutation)
		}
	}
}

func TestUniqPaths(t *testing.T) {
	type testData struct {
		m      int
		n      int
		result int
	}

	datas := []*testData{
		{m: 1, n: 1, result: 1},
		{m: 1, n: 2, result: 1},
		{m: 3, n: 2, result: 3},
		{m: 3, n: 7, result: 28},
		{m: 23, n: 12, result: 193536720},
		{m: 16, n: 16, result: 155117520},
	}

	for _, data := range datas {
		if result := UniquePaths(data.m, data.n); result != data.result {
			t.Errorf("unique_paths error. data:%+v result:%v", data, result)
		}
		if result := UniquePaths(data.n, data.m); result != data.result {
			t.Errorf("unique_paths error. data:%+v result:%v", data, result)
		}
	}
}

func TestUniqPathsWithObstacles(t *testing.T) {
	type testData struct {
		grid  [][]int
		count int
	}

	datas := []*testData{
		{
			grid:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			count: 2,
		},
		{
			grid:  [][]int{{0, 1}, {0, 0}},
			count: 1,
		},
		{
			grid:  [][]int{{0, 1}, {1, 0}},
			count: 0,
		},
		{
			grid: [][]int{
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
				{1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			count: 718991952,
		},
	}

	for _, data := range datas {
		if count := UniquePathsWithObstacles(data.grid); count != data.count {
			t.Errorf("unique_paths_with_obstacles error. data:%+v count:%v", data, count)
		}

	}

}

func TestMiniPathSum(t *testing.T) {
	type testData struct {
		grid    [][]int
		minimum int
	}

	datas := []*testData{
		{
			grid:    [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			minimum: 7,
		},
		{
			grid:    [][]int{{1, 2, 3}, {4, 5, 6}},
			minimum: 12,
		},
		{
			grid:    [][]int{{1}},
			minimum: 1,
		},
	}

	for _, data := range datas {
		if minimum := MinPathSum(data.grid); minimum != data.minimum {
			t.Errorf("min_path_sum error. data:%+v minimum:%v", data, minimum)
		}
	}
}

func TestIsNumber(t *testing.T) {
	type testData struct {
		ss []string
		b  bool
	}

	datas := []*testData{
		{
			ss: []string{"2",
				"0",
				"0089",
				"-0.1",
				"+3.14",
				"4.",
				"-.9",
				".9",
				"2e10",
				"-90E3",
				"3e+7",
				"+6e-1",
				"53.5e93",
				"-123.456e789",
				"46.e3",
			},
			b: true,
		},
		{
			ss: []string{"abc",
				"1a",
				"1e",
				"e3",
				"99e2.5",
				"--6",
				"-+3",
				"95a54e53",
				"e",
				".",
				"4e+",
				"+.",
				".e1",
			},
			b: false,
		},
	}

	for _, data := range datas {
		for _, s := range data.ss {
			if b := IsNumber(s); b != data.b {
				t.Errorf("is_number error. s:%v b:%v data.b:%v", s, b, data.b)
			}
		}
	}
}

func TestPlusOne(t *testing.T) {
	type testData struct {
		digits []int
		result []int
	}

	datas := []*testData{
		{
			digits: []int{1, 2, 3},
			result: []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			result: []int{4, 3, 2, 2},
		},
		{
			digits: []int{9},
			result: []int{1, 0},
		},
		{
			digits: []int{9, 9, 9},
			result: []int{1, 0, 0, 0},
		},
	}

	for _, data := range datas {
		if result := PlusOne(data.digits); !intListEqual(result, data.result) {
			t.Errorf("plus_one error. data:%+v result:%v", data, result)
		}
	}

}

func TestAddBinary(t *testing.T) {
	type testData struct {
		a   string
		b   string
		sum string
	}

	datas := []*testData{
		{
			a:   "11",
			b:   "1",
			sum: "100",
		},
		{
			a:   "1010",
			b:   "1011",
			sum: "10101",
		},
		{
			a:   "1",
			b:   "1",
			sum: "10",
		},
	}

	for _, data := range datas {
		if sum := AddBinary(data.a, data.b); sum != data.sum {
			t.Errorf("add_binary error. data:%+v sum:%v", data, sum)
		}
	}

}

func TestFullJustify(t *testing.T) {
	type testData struct {
		words    []string
		maxWidth int
		result   []string
	}

	datas := []*testData{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			result: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			result: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			result: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, data := range datas {
		if result := FullJustify(data.words, data.maxWidth); !stringListEqual(result, data.result) {
			t.Errorf("full_justify error. data:\n%+v\nresult:\n%+v", strings.Join(data.result, "\n"), strings.Join(result, "\n"))
		}
	}
}

func TestMySqrt(t *testing.T) {
	type testData struct {
		x      int
		result int
	}

	datas := []*testData{
		{
			x:      2,
			result: 1,
		},
		{
			x:      3,
			result: 1,
		},
		{
			x:      4,
			result: 2,
		},
		{
			x:      5,
			result: 2,
		},
		{
			x:      6,
			result: 2,
		},
		{
			x:      7,
			result: 2,
		},
		{
			x:      8,
			result: 2,
		},
		{
			x:      9,
			result: 3,
		},
		{
			x:      10,
			result: 3,
		},
		{
			x:      16,
			result: 4,
		},
	}

	for _, data := range datas {
		if result := MySqrt(data.x); result != data.result {
			t.Errorf("my_sqrt error. data:%+v result:%v", data, result)
		}
	}

}
