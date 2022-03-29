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
