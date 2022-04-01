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

import "testing"

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
