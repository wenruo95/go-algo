package logic

import "testing"

func TestSubsetWithDup(t *testing.T) {
	type testData struct {
		nums   []int
		result [][]int
	}

	datas := []*testData{
		{
			nums: []int{1, 2, 2},
			result: [][]int{
				{},
				{1}, {2},
				{1, 2}, {2, 2},
				{1, 2, 2},
			},
		},
	}
	for _, data := range datas {
		if result := SubsetsWithDup(data.nums); !arraysEqual(result, data.result, true) {
			t.Errorf("subset_with_dup nums:%+v result:%+v expect:%+v", data.nums, jsonstr(result), jsonstr(data.result))
		}
	}

}
