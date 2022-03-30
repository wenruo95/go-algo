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
