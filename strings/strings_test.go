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
	"math"
	"sort"
	"strconv"
	"strings"
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

func TestReverseWords(t *testing.T) {
	type reverseResult struct {
		input  string
		output string
	}

	results := []*reverseResult{
		{
			input:  "the sky is blue",
			output: "blue is sky the",
		},
		{
			input:  "  hello world  ",
			output: "world hello",
		},
		{
			input:  "a good   example",
			output: "example good a",
		},
	}

	for _, result := range results {
		if output := ReverseWords(result.input); output != result.output {
			t.Errorf("reverse_words result:%+v output[%v]:%v expect[%v]:%v",
				result, len(output), output, len(result.output), result.output)
		}
	}

}

func TestPermutations(t *testing.T) {
	type permutationsResult struct {
	}

}

func TestStr2Int(t *testing.T) {
	type str2intResult struct {
		s     string
		i     int
		iserr bool
	}

	results := []*str2intResult{
		// basic
		{
			s:     "100",
			i:     100,
			iserr: false,
		},
		{
			s:     "+100",
			i:     100,
			iserr: false,
		},
		{
			s:     "-100",
			i:     -100,
			iserr: false,
		},
		// invalid char
		{
			s:     "-hello",
			i:     0,
			iserr: true,
		},
		{
			s:     "-520x0",
			i:     0,
			iserr: true,
		},
		//
		{
			s:     strings.Repeat("1", 100),
			i:     0,
			iserr: true,
		},
		{
			s:     "-" + strings.Repeat("1", 100),
			i:     0,
			iserr: true,
		},
		//
		{
			s:     strconv.Itoa(math.MaxInt),
			i:     math.MaxInt,
			iserr: false,
		},
		{
			s:     strconv.Itoa(math.MinInt),
			i:     math.MinInt,
			iserr: false,
		},
		{
			s:     "-" + strconv.Itoa(math.MaxInt),
			i:     -9223372036854775807,
			iserr: false,
		},
		{
			s:     strconv.Itoa(math.MinInt)[1:],
			i:     0,
			iserr: true,
		},
	}
	//t.Logf("min:%v max:%v", math.MinInt, math.MaxInt)

	for _, result := range results {
		i, err := Str2Int(result.s)
		if err != nil && !result.iserr {
			t.Errorf("str2int error match failed. result:%+v error:%v", result, err)
			continue
		}
		if i != result.i {
			t.Errorf("str2int reuslt not matched. result:%+v i:%v ", result, i)
		}
	}

}

func TestReplaceSpace(t *testing.T) {
	type testData struct {
		s      string
		result string
	}

	datas := []*testData{
		{
			s:      "We Are Happy",
			result: "We%20Are%20Happy",
		},
		{
			s:      " ",
			result: "%20",
		},
	}

	for _, data := range datas {
		if result := ReplaceSpace(data.s); result != data.result {
			t.Errorf("replace_space data:%+v result:%v", data, result)
		}
	}

}

func testFindSubString(t *testing.T, fn func(s string, words []string) []int) {
	type testData struct {
		s      string
		words  []string
		indexs []int
	}

	datas := []*testData{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			indexs: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			indexs: []int{},
		},
		{
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			indexs: []int{6, 9, 12},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "good"},
			indexs: []int{8},
		},
	}

	{
		s := strings.Repeat("a", 5000+10)
		words := make([]string, 5000)
		for i := 0; i < 5000; i++ {
			words[i] = "a"
		}
		datas = append(datas, &testData{s: s, words: words, indexs: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
	}

	for _, data := range datas {
		if indexs := fn(data.s, data.words); !intListEqual(indexs, data.indexs) {
			t.Errorf("find_sub_string error. data:%+v indexs:%v", data, indexs)
		}
	}
}

func TestFindSubString(t *testing.T) {
	testFindSubString(t, FindSubstring)
}

func TestFindSubString2(t *testing.T) {
	testFindSubString(t, FindSubstring2)
}

func TestStrCalculate(t *testing.T) {
	type testData struct {
		s   string
		sum float64
	}

	datas := []*testData{
		{
			s:   "1+2+3+4+5+6+7+8",
			sum: 36,
		},
		{
			s:   "12+14-15*16/3",
			sum: -54,
		},
		{
			s:   "12+14-15*16/2*19/20+21+22+23",
			sum: -22,
		},
		{
			s:   "1*2*3*4*5*6",
			sum: 720,
		},
	}

	for _, data := range datas {
		if sum := StrCalculate(data.s); sum != data.sum {
			t.Errorf("strcalculate error. data:%+v sum:%v", data, sum)
		}
	}
}
