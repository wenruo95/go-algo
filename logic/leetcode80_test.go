package logic

import "testing"

func TestRemoveDuplicates2(t *testing.T) {

	type testData struct {
		nums   []int
		k      int
		result []int
	}

	datas := []*testData{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      5,
			result: []int{1, 1, 2, 2, 3},
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			k:      7,
			result: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			nums:   []int{1, 1, 2, 2, 3, 3},
			k:      6,
			result: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, data := range datas {
		if k := RemoveDuplicates2(data.nums); k != data.k || !intListEqual(data.nums[0:k], data.result) {
			t.Errorf("remove_dumplicates error. nums:%v k:%v result:%v", data.nums, k, data.result)
		}
	}

}

func TestSearchInRotateArrayII(t *testing.T) {
	type testData struct {
		nums        []int
		trueTarget  []int
		falseTarget []int
	}

	datas := []*testData{
		{
			nums:        []int{2, 5, 6, 0, 0, 1, 2},
			trueTarget:  []int{0, 1, 2, 5, 6},
			falseTarget: []int{-1, -2, -3, 3, 4, 7, 8, 9, 10},
		},
		{
			nums:        []int{5, 2},
			trueTarget:  []int{2, 5},
			falseTarget: []int{-1, -2, -3, 3, 4, 7, 8, 9, 10},
		},
		{
			nums:        []int{2, 5},
			trueTarget:  []int{2, 5},
			falseTarget: []int{-1, -2, -3, 3, 4, 7, 8, 9, 10},
		},
		{
			nums:        []int{2},
			trueTarget:  []int{2},
			falseTarget: []int{-1, -2, -3, 3, 4, 5, 7, 8, 9, 10},
		},
		{
			nums:        []int{1, 0, 1, 1, 1},
			trueTarget:  []int{0, 1},
			falseTarget: []int{-1, -2, -3, 2, 3, 4},
		},
		{
			nums:        []int{5, 1, 3},
			trueTarget:  []int{5, 3, 1},
			falseTarget: []int{-1, -2, -3, 0, 2, 4, 6, 7, 8},
		},
	}

	for _, data := range datas {
		for _, target := range data.trueTarget {
			if exist := SearchInRotatedSortedArrayII(data.nums, target); !exist {
				t.Errorf("search error. nums:%v target:%v exist:%v expect:true", data.nums, target, exist)
			}
		}
		for _, target := range data.falseTarget {
			if exist := SearchInRotatedSortedArrayII(data.nums, target); exist {
				t.Errorf("search error. nums:%v target:%v exist:%v expect:false", data.nums, target, exist)
			}
		}

	}

}

func TestLargestRectangleArea(t *testing.T) {
	type testData struct {
		heights []int
		area    int
	}

	datas := []*testData{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			area:    10,
		},
		{
			heights: []int{2, 4},
			area:    4,
		},
		{
			heights: []int{1},
			area:    1,
		},
		{
			heights: []int{1, 1},
			area:    2,
		},
		{
			heights: []int{2, 2, 2, 2},
			area:    8,
		},
	}

	for _, data := range datas {
		if area := LargestRectangleArea(data.heights); area != data.area {
			t.Errorf("largest_rectangle_area error. heights:%+v area:%v expect:%v", data.heights, area, data.area)
		}
	}

}

func TestMaximalRectangle(t *testing.T) {
	type testData struct {
		matrix [][]byte
		max    int
	}

	datas := []*testData{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			max: 6,
		},
		{
			matrix: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			max: 0,
		},
		{
			matrix: [][]byte{
				{'0'},
			},
			max: 0,
		},
		{
			matrix: [][]byte{
				{'1'},
			},
			max: 1,
		},
		{
			matrix: [][]byte{
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
			},
			max: 500,
		},
	}

	for _, data := range datas {
		if result := MaximalRectangle(data.matrix); result != data.max {
			t.Errorf("maximal_rectangle matrix:%v result:%v expect:%v", byteArr2Str(data.matrix), result, data.max)
		}
		if result := MaximalRectangle2(data.matrix); result != data.max {
			t.Errorf("maximal_rectangle2 matrix:%v result:%v expect:%v", byteArr2Str(data.matrix), result, data.max)
		}
	}

}
