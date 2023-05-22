package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrapRainWater(t *testing.T) {

	type testData struct {
		heightMap [][]int
		vol       int
	}

	datas := []*testData{
		{
			heightMap: [][]int{
				{1, 4, 3, 1, 3, 2},
				{3, 2, 1, 3, 2, 4},
				{2, 3, 3, 2, 3, 1},
			},
			vol: 4,
		},
		{
			heightMap: [][]int{
				{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3},
			},
			vol: 10,
		},
	}

	for _, data := range datas {

		result := TrapRainWater(data.heightMap)
		assert.Equal(t, data.vol, result)

	}

}
