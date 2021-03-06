/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode50_test.go
*   coder: zemanzeng
*   date : 2022-03-24 22:05:52
*   desc :
*
================================================================*/

package logic

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	type testData struct {
		x      float64
		n      int
		result float64
	}

	datas := []*testData{
		{
			x:      2.00000,
			n:      10,
			result: 1024.00000,
		},
		{
			x:      2.10000,
			n:      3,
			result: 9.26100,
		},
		{
			x:      2.00000,
			n:      -2,
			result: 0.25000,
		},
	}

	for _, data := range datas {
		if result := MyPow(data.x, data.n); result != data.result {
			t.Errorf("my_pow error. data:%+v result:%v", data, result)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	type testData struct {
		n      int
		arrays [][]string
	}

	datas := []*testData{
		{
			n:      1,
			arrays: [][]string{{"Q"}},
		},
		{
			n: 4,
			arrays: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, data := range datas {
		if arrays := SolveNQueens(data.n); !stringArraysEqual(arrays, data.arrays, false) {
			t.Errorf("solve_n_queens error. data:%+v arrays:%+v", data, arrays)
		}
	}

}

func TestTotalNQueens(t *testing.T) {
	type testData struct {
		n     int
		total int
	}

	datas := []*testData{
		{n: 1, total: 1},
		{n: 4, total: 2},
	}

	for _, data := range datas {
		if total := TotalNQueens(data.n); total != data.total {
			t.Errorf("total_n_queens error. data:%+v total:%v", data, total)
		}
	}
}

func testMaxSubArray(t *testing.T, fn func([]int) int) {
	type testData struct {
		nums []int
		sum  int
	}

	datas := []*testData{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			sum:  6,
		},
		{
			nums: []int{1},
			sum:  1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			sum:  23,
		},
		{
			nums: []int{-1},
			sum:  -1,
		},
	}

	for _, data := range datas {
		if sum := fn(data.nums); sum != data.sum {
			t.Errorf("max_sub_array error. data:%+v sum:%v", data, sum)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	testMaxSubArray(t, MaxSubArray)
}

func TestMaxSubArray2(t *testing.T) {
	testMaxSubArray(t, MaxSubArray2)
}

func TestSpiralOrder(t *testing.T) {
	type testData struct {
		matrix [][]int
		list   []int
	}

	datas := []*testData{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			list:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			list:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{{1}},
			list:   []int{1},
		},
		{
			matrix: [][]int{{6, 9, 7}},
			list:   []int{6, 9, 7},
		},
		{
			matrix: [][]int{{7}, {9}, {6}},
			list:   []int{7, 9, 6},
		},
	}

	for _, data := range datas {
		if list := SpiralOrder(data.matrix); !intListEqual(list, data.list) {
			t.Errorf("spiral_order error. data:%+v list:%v", data, list)

		}

	}

}

func TestCanJump(t *testing.T) {
	type testData struct {
		nums []int
		jump bool
	}

	datas := []*testData{
		{
			nums: []int{2, 3, 1, 1, 4},
			jump: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4},
			jump: false,
		},
		{
			nums: []int{1, 2},
			jump: true,
		},
	}

	for _, data := range datas {
		if jump := CanJump(data.nums); jump != data.jump {
			t.Errorf("can_jump error. data:%+v jump:%v", data, jump)
		}
	}
}

func TestMergeIntervals(t *testing.T) {
	type testData struct {
		intervals [][]int
		arrays    [][]int
	}

	datas := []*testData{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			arrays:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			arrays:    [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {5, 6}},
			arrays:    [][]int{{1, 4}, {5, 6}},
		},
	}

	for _, data := range datas {
		arrays := MergeIntervals(data.intervals)
		if !arraysEqual(arrays, data.arrays, false) {
			t.Errorf("merge_intervals error. data:%+v arrays:%v", data, arrays)

		}
	}

}

func TestMergeInsertInterval(t *testing.T) {
	type testData struct {
		intervals   [][]int
		newInterval []int
		arrays      [][]int
	}

	datas := []*testData{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			arrays:      [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			arrays:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{{4, 6}, {7, 8}},
			newInterval: []int{1, 3},
			arrays:      [][]int{{1, 3}, {4, 6}, {7, 8}},
		},
		{
			intervals:   [][]int{{4, 6}, {7, 8}},
			newInterval: []int{9, 10},
			arrays:      [][]int{{4, 6}, {7, 8}, {9, 10}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			arrays:      [][]int{{1, 5}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{5, 7},
			arrays:      [][]int{{1, 7}},
		},
		{
			intervals:   [][]int{{5, 7}},
			newInterval: []int{1, 5},
			arrays:      [][]int{{1, 7}},
		},
		{
			intervals:   [][]int{{4, 6}, {7, 8}},
			newInterval: []int{6, 7},
			arrays:      [][]int{{4, 8}},
		},
		{
			intervals:   [][]int{{4, 6}, {7, 8}, {10, 11}},
			newInterval: []int{6, 7},
			arrays:      [][]int{{4, 8}, {10, 11}},
		},
	}

	for _, data := range datas {
		arrays := MergeInsertInterval(data.intervals, data.newInterval)
		if !arraysEqual(arrays, data.arrays, false) {
			t.Errorf("merge_insert_intervals error. data:%+v arrays:%v", data, arrays)

		}
	}

}

func TestLengthOfLastWord(t *testing.T) {
	type testData struct {
		s      string
		length int
	}

	datas := []*testData{
		{
			s:      "Hello World",
			length: 5,
		},
		{
			s:      "   fly me   to   the moon  ",
			length: 4,
		},
		{
			s:      "luffy is still joyboy",
			length: 6,
		},
		{
			s:      "Hello",
			length: 5,
		},
	}

	for _, data := range datas {
		if length := LengthOfLastWord(data.s); length != data.length {
			t.Errorf("length_of_last_word error. data:%+v length:%v", data, length)
		}
	}

}

func TestGenerateMatrix(t *testing.T) {
	type testData struct {
		n      int
		matrix [][]int
	}

	datas := []*testData{
		{
			n:      2,
			matrix: [][]int{{1, 2}, {4, 3}},
		},
		{
			n:      3,
			matrix: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			n:      4,
			matrix: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		},
	}

	for _, data := range datas {
		if arrays := GenerateMatrix(data.n); !arraysEqual(arrays, data.matrix, false) {
			t.Errorf("generate_matrix error. data:%+v arrays:%+v", data, arrays)
		}
	}

}
