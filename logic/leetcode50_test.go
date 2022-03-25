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
