package logic

import "testing"

func TestContainDuplicate(t *testing.T) {
	type testData struct {
		nums   []int
		result bool
	}

	datas := []*testData{
		{nums: []int{1, 2, 3, 1}, result: true},
		{nums: []int{1, 2, 3, 4}, result: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, result: true},
	}
	for _, data := range datas {
		if result := ContainsDuplicate(data.nums); result != data.result {
			t.Errorf("contain_duplicate nums:%v result:%v expect:%v", data.nums, result, data.result)
		}
	}
}
