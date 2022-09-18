package logic

import "testing"

func TestProductExceptSelf(t *testing.T) {
	type testData struct {
		nums   []int
		result []int
	}

	datas := []*testData{
		{
			nums:   []int{1, 2, 3, 4},
			result: []int{24, 12, 8, 6},
		},
		{
			nums:   []int{-1, 1, 0, -3, 3},
			result: []int{0, 0, 9, 0, 0},
		},
	}
	for _, data := range datas {
		if result := ProductExceptSelf(data.nums); !intListEqual(result, data.result) {
			t.Errorf("product_except_self nums:%+v result:%+v expect:%+v", data.nums, result, data.result)
		}
	}
}
