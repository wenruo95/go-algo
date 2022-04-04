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
