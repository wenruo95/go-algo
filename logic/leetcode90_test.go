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

func TestNumDecodings(t *testing.T) {
	type testData struct {
		s string
		n int
	}

	datas := []*testData{
		{s: "12", n: 2},
		{s: "226", n: 3},
		{s: "06", n: 0},
		{s: "0", n: 0},
		{s: "6", n: 1},
	}

	for _, data := range datas {
		if n := NumDecodings(data.s); n != data.n {
			t.Errorf("num_decodings s:%v n:%v expect:%v", data.s, n, data.n)
		}
	}
}
