/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : sum_test.go
*   coder: zemanzeng
*   date : 2022-02-02 21:01:54
*   desc : sum test case
*
================================================================*/

package logic

import "testing"

func TestTreeSumClosest(t *testing.T) {

	type closestResult struct {
		nums   []int
		target int
		output int
	}

	results := []*closestResult{
		{

			nums:   []int{-1, 2, 1, -4},
			target: 1,
			output: 2,
		},
		{

			nums:   []int{0, 0, 0},
			target: 1,
			output: 0,
		},
	}

	for _, result := range results {
		output := ThreeSumClosest(result.nums, result.target)
		if output != result.output {
			t.Errorf("three_sum_closest result:%+v output:%v", result, output)
		}
	}

}
