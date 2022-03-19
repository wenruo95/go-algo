/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode330_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:42:27
*   desc :
*
================================================================*/

package logic

import "testing"

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
