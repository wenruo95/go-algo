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

import "testing"

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
