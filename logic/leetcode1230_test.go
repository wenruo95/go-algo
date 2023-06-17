package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStraightLine(t *testing.T) {
	type testData struct {
		coordinates [][]int
		res         bool
	}

	datas := []*testData{
		{
			coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			res:         true,
		},
		{
			coordinates: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			res:         false,
		},
		{
			coordinates: [][]int{{0, 0}, {0, 1}, {0, -1}},
			res:         true,
		},
		{
			coordinates: [][]int{{0, 0}, {1, 0}, {-1, 0}},
			res:         true,
		},
	}
	for _, data := range datas {
		res := CheckStraightLine(data.coordinates)
		assert.Equal(t, data.res, res)
	}

}
