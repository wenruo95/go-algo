package logic

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	type testData struct {
		nums []int
		k    int
		res  []int
	}

	datas := []*testData{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, res: []int{1, 2}},
		{nums: []int{1, 1, 1, 2, 2, 2, 3, 4, 4, 4, 5, 5, 5}, k: 4, res: []int{1, 2, 4, 5}},
		{nums: []int{1}, k: 1, res: []int{1}},
	}

	for _, data := range datas {
		res := TopKFrequent(data.nums, data.k)
		sort.Ints(data.res)
		sort.Ints(res)
		assert.Equal(t, data.res, res)
	}

}
