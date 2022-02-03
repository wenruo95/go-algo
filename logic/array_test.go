/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : array_test.go
*   coder: zemanzeng
*   date : 2022-02-03 19:50:24
*   desc :
*
================================================================*/

package logic

import "testing"

func TestRemoveDumplicates(t *testing.T) {
	type removeResult struct {
		input  []int
		output []int
		k      int
	}

	results := []*removeResult{
		{
			input:  []int{1, 1, 2},
			output: []int{1, 2},
			k:      2,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			k:      5,
			output: []int{0, 1, 2, 3, 4},
		},
	}

	for _, result := range results {
		k := RemoveDuplicates(result.input)
		if k != result.k {
			t.Errorf("remove_dumplicates result:%+v k:%v", result, k)
		}
		for i := 0; i < k; i++ {
			if result.output[i] != result.input[i] {
				t.Errorf("remove_dumplicates result:%+v k:%v index:%v", result, k, i)
			}
		}
	}

}
