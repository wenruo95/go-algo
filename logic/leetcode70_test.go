/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode70_test.go
*   coder: zemanzeng
*   date : 2022-04-03 22:12:00
*   desc :
*
================================================================*/

package logic

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	type testData struct {
		n    int
		step int
	}

	datas := []*testData{
		{n: 0, step: 0},
		{n: 1, step: 1},
		{n: 2, step: 2},
		{n: 3, step: 3},
		{n: 4, step: 5},
		{n: 5, step: 8},
		{n: 6, step: 13},
	}

	for _, data := range datas {
		if step := ClimbStairs(data.n); step != data.step {
			t.Errorf("climb_stairs error. data:%+v step:%v", data, step)
		}
	}
}

func TestSimplifyPath(t *testing.T) {
	type testData struct {
		path   string
		result string
	}

	datas := []*testData{
		{
			path:   "/home/",
			result: "/home",
		},
		{
			path:   "/../",
			result: "/",
		},
		{
			path:   "/home//foo/",
			result: "/home/foo",
		},
		{
			path:   "/a/./b/../../c/",
			result: "/c",
		},
		{
			path:   "/a//b////c/d//././/..",
			result: "/a/b/c",
		},
		{
			path:   "/.././GVzvE/./xBjU///../..///././//////T/../../.././zu/q/e",
			result: "/zu/q/e",
		},
	}

	for _, data := range datas {
		if result := SimplifyPath(data.path); result != data.result {
			t.Errorf("simplify_path error. data:%+v result:%v", data, result)
		}
	}

}

func TestMinDistance(t *testing.T) {
	type testData struct {
		word1    string
		word2    string
		distance int
	}

	datas := []*testData{
		{
			word1:    "horse",
			word2:    "ros",
			distance: 3,
		},
		{
			word1:    "intention",
			word2:    "execution",
			distance: 5,
		},
	}

	for _, data := range datas {
		if distance := MinDistance(data.word1, data.word2); distance != data.distance {
			t.Errorf("min_distance error. data:%+v distance:%v", data, distance)
		}
	}
}

func TestSetMatrixZeroes(t *testing.T) {
	type testData struct {
		matrix [][]int
		arrays [][]int
	}

	datas := []*testData{
		{
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			arrays: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			arrays: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, data := range datas {
		SetMatrixZeroes(data.matrix)
		if !arraysEqual(data.matrix, data.arrays, false) {
			t.Errorf("set_matrix_zeroes error. matrix:%+v arrays:%+v", data.matrix, data.arrays)
		}
	}

}

func TestSearchMatrix(t *testing.T) {
	type testData struct {
		matrix  [][]int
		targets []int
		find    bool
	}

	datas := []*testData{
		{
			matrix:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			targets: []int{1, 3, 5, 7, 10, 11, 16, 20, 23, 30, 34, 60},
			find:    true,
		},

		{
			matrix:  [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			targets: []int{2, 4, 6, 8, 9, 12, 13, 14, 15, 17, 18, 19, 21, 22, 25, 26, 27, 28, 29, 31, 32, 33, 35},
			find:    false,
		},
	}

	for _, data := range datas {
		for _, target := range data.targets {
			if find := SearchMatrix(data.matrix, target); find != data.find {
				t.Errorf("search_matrix error. data:%+v target:%v find:%v expect:%v", data.matrix, target, find, data.find)
			}
		}

	}
}

func TestSortColors(t *testing.T) {
	type testData struct {
		nums  []int
		array []int
	}

	datas := []*testData{
		{
			nums:  []int{2, 0, 2, 1, 1, 0},
			array: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums:  []int{2, 0, 1},
			array: []int{0, 1, 2},
		},
	}

	for _, data := range datas {
		if SortColors(data.nums); !intListEqual(data.nums, data.array) {
			t.Errorf("sort_colors error. data:%+v", data)
		}
	}

}

func TestMinWindow(t *testing.T) {
	type testData struct {
		s      string
		t      string
		window string
	}

	datas := []*testData{
		{s: "ADOBECODEBANC", t: "ABC", window: "BANC"},
		{s: "a", t: "a", window: "a"},
		{s: "a", t: "aa", window: ""},
		{s: "aa", t: "aa", window: "aa"},
		{s: "ab", t: "b", window: "b"},
		{s: "aaaaaaaaaaaabbbbbcdd", t: "abcdd", window: "abbbbbcdd"},
	}

	for _, data := range datas {
		if window := MinWindow(data.s, data.t); window != data.window {
			t.Errorf("min_window error. data:%+v window:%v", data, window)
		}
		if window := MinWindow2(data.s, data.t); window != data.window {
			t.Errorf("min_window2 error. data:%+v window:%v", data, window)
		}
		if window := MinWindow3(data.s, data.t); window != data.window {
			t.Errorf("min_window3 error. data:%+v window:%v", data, window)
		}

	}

}
