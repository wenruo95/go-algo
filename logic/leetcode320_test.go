/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode320_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:45:21
*   desc :
*
================================================================*/

package logic

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	type longestPathResult struct {
		matrix [][]int
		output int
	}

	results := []*longestPathResult{
		{
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			output: 4,
		},
		{
			matrix: [][]int{
				{6, 8},
				{7, 2},
			},
			output: 2,
		},
	}
	for _, result := range results {
		if output := LongestIncreasingPath(result.matrix); output != result.output {
			t.Errorf("longest_increasing_path result:%+v output:%v", result, output)
		}
	}

}
