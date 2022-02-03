/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic_test.go
*   coder: zemanzeng
*   date : 2022-02-03 01:16:43
*   desc :
*
================================================================*/

package logic

import (
	"sort"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	type parenthesesResult struct {
		s     string
		valid bool
	}
	results := []*parenthesesResult{
		{
			s:     "()",
			valid: true,
		},
		{
			s:     "()[]{}",
			valid: true,
		},
		{
			s:     "(]",
			valid: false,
		},
	}

	for _, result := range results {
		if valid := IsValidParentheses(result.s); valid != result.valid {
			t.Errorf("is_valid_parentheses result:%+v valid:%v", result, valid)
		}
	}

}

func stringListEquals(dst, src []string) bool {
	if len(dst) != len(src) {
		return false
	}

	sort.Strings(dst)
	sort.Strings(src)
	for i := 0; i < len(dst); i++ {
		if dst[i] != src[i] {
			return false
		}
	}

	return true
}

func TestGenerateParenthesis(t *testing.T) {
	type generateParenthesis struct {
		n      int
		output []string
	}

	results := []*generateParenthesis{
		{
			n:      1,
			output: []string{"()"},
		},
		{
			n:      2,
			output: []string{"()()", "(())"},
		},
		{
			n:      3,
			output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n: 4,
			output: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())",
				"(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	for _, result := range results {
		output := GenerateParenthesis(result.n)
		if !stringListEquals(output, result.output) {
			t.Errorf("generate_parenthesis result:%+v output:%+v", result, output)
		}

	}

}

func TestDivide(t *testing.T) {
	type divideResult struct {
		dividend int
		divisor  int
		quotient int
	}

	results := []*divideResult{
		{
			dividend: 10,
			divisor:  3,
			quotient: 3,
		},
		{
			dividend: 7,
			divisor:  -3,
			quotient: -2,
		},
		{
			dividend: 1,
			divisor:  2,
			quotient: 0,
		},
		{
			dividend: -2147483648,
			divisor:  -1,
			quotient: 2147483647,
		},
	}

	for _, result := range results {
		quotient := Divide(result.dividend, result.divisor)
		if quotient != result.quotient {
			t.Errorf("divide result:%+v quotient:%v", result, quotient)
		}
	}

}

func TestSearch(t *testing.T) {
	type searchResult struct {
		nums   []int
		target int
		index  int
	}

	results := []*searchResult{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			index:  4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			index:  -1,
		},
		{
			nums:   []int{1},
			target: 0,
			index:  -1,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 4,
			index:  3,
		},
		{
			nums:   []int{3, 1},
			target: 1,
			index:  1,
		},
	}
	for _, result := range results {
		index := SearchInRotatedSortedArray(result.nums, result.target)
		if index != result.index {
			t.Errorf("search_in_rotated_sorted_array result:%+v index:%v", result, index)
		}
	}

}
