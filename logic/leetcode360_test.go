/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode360_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:44:40
*   desc :
*
================================================================*/

package logic

import "testing"

// leetcode 365
func TestCanMeasureWater(t *testing.T) {
	type measureResult struct {
		jug1   int
		jug2   int
		target int
		output bool
	}

	results := []*measureResult{
		{jug1: 3, jug2: 5, target: 4, output: true},
		{jug1: 2, jug2: 6, target: 5, output: false},
		{jug1: 1, jug2: 2, target: 3, output: true},
		{jug1: 9, jug2: 6, target: 1, output: false},
		{jug1: 10000, jug2: 10001, target: 1, output: true},
		{jug1: 4, jug2: 6, target: 8, output: true},
		{jug1: 34, jug2: 5, target: 6, output: true},
		{jug1: 1, jug2: 1, target: 12, output: false},
	}

	for _, result := range results {
		output := CanMeasureWater(result.jug1, result.jug2, result.target)
		if output != result.output {
			t.Errorf("can_measure_water result:%+v output:%v", result, output)
		}
	}

}
