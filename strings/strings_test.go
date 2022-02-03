/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : strings_test.go
*   coder: zemanzeng
*   date : 2022-02-02 15:07:03
*   desc : strings test case
*
================================================================*/

package strings

import (
	"sort"
	"testing"
)

func intListEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

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

func stringListEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestLetterCombinations(t *testing.T) {

	type letterResult struct {
		digits string
		output []string
	}

	results := []*letterResult{
		{
			digits: "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			output: []string{},
		},
		{
			digits: "2",
			output: []string{"a", "b", "c"},
		},
	}

	for _, result := range results {
		combs := LetterCombinations(result.digits)
		if !stringListEqual(combs, result.output) {
			t.Errorf("letter_combineations. digits:%+v output:%+v combs:%+v",
				result.digits, result.output, combs)
		}

	}

}

func TestRegularIsMatch(t *testing.T) {
	type regularResult struct {
		s     string
		p     string
		match bool
	}

	results := []*regularResult{
		{
			s:     "aa",
			p:     "a",
			match: false,
			// Explanation: "a" does not match the entire string "aa".
		},
		{
			s:     "aa",
			p:     "a*",
			match: true,
			// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
		},
		{
			s:     "ab",
			p:     ".*",
			match: true,
			// Explanation: ".*" means "zero or more (*) of any character (.)".
		},
		{
			s:     "aab",
			p:     "c*a*b",
			match: true,
		},
		{
			s:     "mississippi",
			p:     "mis*is*p*.",
			match: false,
		},
		{
			s:     "aaa",
			p:     "ab*a*c*a",
			match: true,
		},
	}

	for _, result := range results {
		if match := RegularIsMatch(result.s, result.p); match != result.match {
			t.Errorf("regular_is_match result:%+v match:%v", result, match)
		}
	}
}

func TestStrStr(t *testing.T) {
	type strResult struct {
		haystack string
		needle   string
		index    int
	}

	results := []*strResult{
		{
			haystack: "hello",
			needle:   "ll",
			index:    2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			index:    -1,
		},
		{
			haystack: "",
			needle:   "",
			index:    0,
		},
		{
			haystack: "aaa",
			needle:   "aaaa",
			index:    -1,
		},
	}

	for _, result := range results {
		if index := StrStr(result.haystack, result.needle); index != result.index {
			t.Errorf("str_str result:%+v index:%v", result, index)
		}
	}

}
