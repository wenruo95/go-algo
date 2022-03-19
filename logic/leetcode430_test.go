/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode430_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:29:46
*   desc :
*
================================================================*/

package logic

import "testing"

// leetcode 438
func TestFindAnagrams(t *testing.T) {
	type anagramsResult struct {
		s      string
		p      string
		output []int
	}

	results := []*anagramsResult{
		{
			s:      "cbaebabacd",
			p:      "abc",
			output: []int{0, 6},
		},
		{
			s:      "abab",
			p:      "ab",
			output: []int{0, 1, 2},
		},
		{
			s:      "baa",
			p:      "aa",
			output: []int{1},
		},
	}

	for _, result := range results {
		indexs := FindAnagrams(result.s, result.p)
		if !intListEqual(indexs, result.output) {
			t.Errorf("find_anagrams error. s:%v p:%v output:%+v result:%+v",
				result.s, result.p, result.output, indexs)
		}
	}
}
