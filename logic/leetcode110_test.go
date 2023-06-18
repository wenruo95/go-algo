package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {

	type testData struct {
		list []interface{}
		res  bool
	}

	datas := []*testData{
		{
			list: []interface{}{3, 9, 20, nil, nil, 15, 7},
			res:  true,
		},
		{
			list: []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			res:  false,
		},
		{
			list: []interface{}{},
			res:  true,
		},
		{
			list: []interface{}{1, 2, 3, 4, 5, 6, nil, 8},
			res:  true,
		},
	}
	for _, data := range datas {
		node := genTreeNode(data.list)
		res := IsBalanced(node)
		assert.Equal(t, data.res, res)
	}

}

func TestMinDepth(t *testing.T) {

	type testData struct {
		list  []interface{}
		depth int
	}

	datas := []*testData{
		{
			list:  []interface{}{3, 9, 20, nil, nil, 15, 7},
			depth: 2,
		},
		{
			list:  []interface{}{1, nil, 2},
			depth: 2,
		},
		{
			list: []interface{}{
				2,
				nil, 3,
				nil, nil, nil, 4,
				nil, nil, nil, nil, nil, nil, nil, 5,
				nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil, nil, 6,
			},
			depth: 5,
		},
	}

	for _, data := range datas {
		node := genTreeNode(data.list)
		depth := MinDepth(node)
		assert.Equal(t, data.depth, depth)
		t.Logf("list:%+v expect:%v result:%v", data.list, data.depth, depth)
	}

}
