/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode20_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:18:12
*   desc :
*
================================================================*/

package logic

import "testing"

// leetcode 20
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

// leetcode 22
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
		if !stringListItemEqual(output, result.output) {
			t.Errorf("generate_parenthesis result:%+v output:%+v", result, output)
		}

	}

}

// leetcode 26
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

// leetcode 27
func TestRemoveElement(t *testing.T) {
	type testData struct {
		nums  []int
		val   int
		count int
	}

	datas := []*testData{
		{
			nums:  []int{3, 2, 2, 3},
			val:   3,
			count: 2,
		},
		{
			nums:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:   2,
			count: 5,
		},
	}

	for _, data := range datas {
		if count := RemoveElement(data.nums, data.val); count != data.count {
			t.Errorf("remove_element data:%+v count:%v", data, count)
		}
	}

}

// leetcode 28
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

// leetcode 29
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
