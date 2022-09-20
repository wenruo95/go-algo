package logic

import (
	"math"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	type testData struct {
		num   uint32
		count int
	}

	datas := []*testData{
		{num: 11, count: 3},
		{num: 128, count: 1},
		{num: math.MaxUint32 - 2, count: 31},
	}
	for _, data := range datas {
		if count := HammingWeight(data.num); count != data.count {
			t.Errorf("hamming_weight num(%b):%v count:%v expect:%v", data.num, data.num, count, data.count)
		}
	}

}
