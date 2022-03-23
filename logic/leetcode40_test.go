/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode40_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:33:50
*   desc :
*
================================================================*/

package logic

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type testData struct {
		candidates []int
		target     int
		arrays     [][]int
	}

	datas := []*testData{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			arrays: [][]int{
				{2, 2, 3}, {7},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			arrays: [][]int{
				{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
			},
		},
		{
			candidates: []int{2},
			target:     1,
			arrays:     [][]int{},
		},
		{
			candidates: []int{3, 12, 9, 11, 6, 7, 8, 5, 4},
			target:     15,
			arrays: [][]int{
				{3, 3, 3, 3, 3},
				{3, 3, 3, 6},
				{3, 3, 4, 5},
				{3, 3, 9},
				{3, 4, 4, 4},
				{3, 4, 8},
				{3, 5, 7},
				{3, 6, 6},
				{3, 12},
				{4, 4, 7},
				{4, 5, 6},
				{4, 11},
				{5, 5, 5},
				{6, 9},
				{7, 8},
			},
		},
	}

	for _, data := range datas {
		if arrays := CombinationSum(data.candidates, data.target); !arraysEqual(arrays, data.arrays, true) {
			t.Errorf("combination_sum error. data:%+v arrays:%v", data, arrays)
		}
	}

}

func TestCombinationSum2(t *testing.T) {
	type testData struct {
		candidates []int
		target     int
		arrays     [][]int
	}

	datas := []*testData{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			arrays: [][]int{
				{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
			},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			arrays: [][]int{
				{1, 2, 2}, {5},
			},
		},
		{
			candidates: []int{4, 1, 1, 4, 4, 4, 4, 2, 3, 5},
			target:     10,
			arrays: [][]int{
				{1, 1, 3, 5}, {1, 1, 4, 4}, {1, 2, 3, 4}, {1, 4, 5}, {2, 3, 5}, {2, 4, 4},
			},
		},
	}

	for _, data := range datas {
		if arrays := CombinationSum2(data.candidates, data.target); !arraysEqual(arrays, data.arrays, true) {
			t.Errorf("combination_sum_2 error. data:%+v arrays:%v", data, arrays)
		}
	}

}
func testMultiply(t *testing.T, fn func(string, string) string) {

	type testData struct {
		s1     string
		s2     string
		result string
	}

	datas := []*testData{
		{s1: "1", s2: "2", result: "2"},
		{s1: "0", s2: "2", result: "0"},
		{s1: "9", s2: "2", result: "18"},
		{s1: "10", s2: "2", result: "20"},
		{s1: "10", s2: "10", result: "100"},
		{s1: "20", s2: "50", result: "1000"},
		{s1: "234", s2: "456", result: "106704"},
		{s1: "123", s2: "456", result: "56088"},
		{s1: "0", s2: "9133", result: "0"},
		{s1: "498828660196", s2: "840477629533", result: "419254329864656431168468"},
	}

	for _, data := range datas {
		if result := fn(data.s1, data.s2); result != data.result {
			t.Errorf("multiply error. data:%+v result:%v", data, result)
		}
		if result := fn(data.s2, data.s1); result != data.result {
			t.Errorf("multiply error. data:%+v result:%v", data, result)
		}
	}

}

func testFirstMissingPositive(t *testing.T, fn func([]int) int) {
	type testData struct {
		nums    []int
		missing int
	}

	datas := []*testData{
		{
			nums:    []int{1, 2, 0},
			missing: 3,
		},
		{
			nums:    []int{3, 4, -1, 1},
			missing: 2,
		},
		{
			nums:    []int{7, 8, 9, 11, 12},
			missing: 1,
		},
		{
			nums:    []int{-5},
			missing: 1,
		},
		{
			nums:    []int{-5, -3},
			missing: 1,
		},
		{
			nums:    []int{1},
			missing: 2,
		},
	}

	for _, data := range datas {
		if missing := fn(data.nums); missing != data.missing {
			t.Errorf("first_missing_positive error. data:%+v missing:%v", data, missing)
		}
	}

}

// leetcode 41
func TestFirstMissingPositive(t *testing.T) {
	testFirstMissingPositive(t, FirstMissingPositive)
}

func TestFirstMissingPositive2(t *testing.T) {
	testFirstMissingPositive(t, FirstMissingPositive2)
}

// leetcode 42
func TestTrap(t *testing.T) {

	type testData struct {
		height []int
		sum    int
	}

	datas := []*testData{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			sum:    6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			sum:    9,
		},
	}

	for _, data := range datas {
		if sum := Trap(data.height); sum != data.sum {
			t.Errorf("trap error. data:%+v sum:%v", data, sum)
		}
	}
}

// leetcode 43
func TestMultiply(t *testing.T) {
	testMultiply(t, Multiply)
}

func TestMultiply2(t *testing.T) {
	testMultiply(t, Multiply2)
}

func TestIsMatch(t *testing.T) {
	type testData struct {
		s string
		p string
		b bool
	}

	datas := []*testData{
		{
			s: "aa",
			p: "a",
			b: false,
		},
		{
			s: "aa",
			p: "*",
			b: true,
		},
		{
			s: "cb",
			p: "?a",
			b: false,
		},
		{
			s: "adceb",
			p: "*a*b",
			b: true,
		},
		{
			s: "aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba",
			p: "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*",
			b: true,
		},
		{
			s: "acdcb",
			p: "a*c?b",
			b: false,
		},
		{
			s: "a",
			p: "",
			b: false,
		},
	}

	for _, data := range datas {
		if b := IsMatch(data.s, data.p); b != data.b {
			t.Errorf("is_match error. data:%+v b:%v", data, b)
		}

	}

}

func testJumpGameII(t *testing.T, fn func([]int) int) {
	type testData struct {
		nums []int
		step int
	}

	datas := []*testData{
		{
			nums: []int{2, 3, 1, 1, 4},
			step: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			step: 2,
		},
		{
			nums: []int{2},
			step: 0,
		},
		{
			nums: []int{2, 4},
			step: 1,
		},
	}

	for _, data := range datas {
		if step := fn(data.nums); step != data.step {
			t.Errorf("jump_game_ii error. data:%+v step:%v", data, step)
		}
	}
}

func TestJumpGameII(t *testing.T) {
	testJumpGameII(t, JumpGameII)
}

func TestJumpGameII2(t *testing.T) {
	testJumpGameII(t, JumpGameII2)
}

func TestPermute(t *testing.T) {
	type testData struct {
		nums   []int
		arrays [][]int
	}
	datas := []*testData{
		{
			nums: []int{1, 2, 3},
			arrays: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			nums:   []int{1},
			arrays: [][]int{{1}},
		},
	}

	for _, data := range datas {
		if arrays := Permute(data.nums); !arraysEqual(arrays, data.arrays, false) {
			t.Errorf("permute error. data:%+v arrays:%v", data, arrays)
		}
	}

}

func TestPermuteUnique(t *testing.T) {
	type testData struct {
		nums   []int
		arrays [][]int
	}
	datas := []*testData{
		{
			nums: []int{1, 2, 3},
			arrays: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			nums:   []int{1},
			arrays: [][]int{{1}},
		},
		{
			nums: []int{1, 2, 1},
			arrays: [][]int{
				{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
			},
		},
		{
			nums: []int{2, 2, 1, 1},
			arrays: [][]int{
				{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1},
			},
		},
		{
			nums: []int{0, 1, 0, 0, 9},
			arrays: [][]int{
				{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0}, {0, 0, 9, 0, 1}, {0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0}, {0, 1, 9, 0, 0}, {0, 9, 0, 0, 1}, {0, 9, 0, 1, 0}, {0, 9, 1, 0, 0},
				{1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0}, {1, 9, 0, 0, 0},
				{9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0},
			},
		},
	}

	for _, data := range datas {
		if arrays := PermuteUnique(data.nums); !arraysEqual(arrays, data.arrays, false) {
			t.Errorf("permute_unique error. data:%+v arrays:%v", data, arrays)
		}
	}

}

func TestRotate(t *testing.T) {
	type testData struct {
		matrix [][]int
		result [][]int
	}

	datas := []*testData{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			result: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			result: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, data := range datas {
		Rotate(data.matrix)
		if !arraysEqual(data.matrix, data.result, false) {
			t.Errorf("rotate error. matrix:%+v result:%v", data.matrix, data.result)
		}
	}
}
