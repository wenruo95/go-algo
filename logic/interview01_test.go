/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : interview01_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:32:31
*   desc :
*
================================================================*/

package logic

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestMaxSplitN(t *testing.T) {

	type testData struct {
		n   int
		max int
	}

	datas := []*testData{
		{
			n:   0,
			max: 0,
		},
		{
			n:   10,
			max: 36,
		},
		{
			n:   20,
			max: 1458,
		},
	}

	for _, data := range datas {
		if max := GetMaxSplitN(data.n); max != data.max {
			t.Errorf("get_max_split_n data:%+v max:%v", data, max)
		}
	}

}

func TestStrPermutations(t *testing.T) {
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

func TestFindMaxSeq(t *testing.T) {
	type testData struct {
		c      int
		sizes  []int
		result []int
	}

	datas := []*testData{
		{
			c:      1,
			sizes:  []int{1, 2, 3, 5, 4},
			result: []int{0, 0},
		},
		{
			c:      7,
			sizes:  []int{1, 2, 3, 5, 4},
			result: []int{0, 2},
		},
		{
			c:      7,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{5, 5},
		},
		{
			c:      21,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{1, 5},
		},
		{
			c:      22,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{0, 5},
		},
		{
			c:      23,
			sizes:  []int{1, 2, 3, 5, 4, 7},
			result: []int{0, 5},
		},
	}

	for _, data := range datas {
		result := FindMaxSeq(data.c, data.sizes)
		if len(result) != 2 ||
			result[0] != data.result[0] || result[1] != data.result[1] {
			t.Errorf("find_max_seq error. data:%+v result:%v", data, result)
		}

	}
}

func TestMaxSplitStringN(t *testing.T) {

	type testData struct {
		s      string
		output []string
	}

	datas := []*testData{
		{
			s:      "ababc",
			output: []string{"abab", "c"},
		},
	}

	for _, data := range datas {
		output := MaxSplitStringN(data.s)
		if !stringListEqual(output, data.output) {
			t.Errorf("max_split_string_n data:%+v output:%v", data, output)
		}
	}

}
