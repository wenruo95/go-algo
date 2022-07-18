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
