package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	type testData struct {
		p    []interface{}
		q    []interface{}
		same bool
	}

	datas := []*testData{
		{
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 3},
			same: true,
		},
		{
			p:    []interface{}{1, 2},
			q:    []interface{}{1, nil, 2},
			same: false,
		},

		{
			p:    []interface{}{1, 2, 1},
			q:    []interface{}{1, 1, 2},
			same: false,
		},
	}

	for _, data := range datas {
		p := genTreeNode(data.p)
		q := genTreeNode(data.q)
		same := IsSameTree(p, q)
		assert.Equal(t, data.same, same)
	}
}

func TestIsSymmetric(t *testing.T) {
	type testData struct {
		list      []interface{}
		symmetric bool
	}

	datas := []*testData{
		{
			list:      []interface{}{1, 2, 2, 3, 4, 4, 3},
			symmetric: true,
		},
		{
			list:      []interface{}{1, 2, 2, nil, 3, nil, 3},
			symmetric: false,
		},
	}

	for _, data := range datas {
		list := genTreeNode(data.list)
		symmetric := IsSymmetric(list)
		assert.Equal(t, data.symmetric, symmetric)
	}
}

func TestLevelOrder(t *testing.T) {

	type testData struct {
		list []interface{}
		res  [][]int
	}

	datas := []*testData{
		{
			list: []interface{}{3, 9, 20, nil, nil, 15, 7},
			res:  [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			list: []interface{}{1},
			res:  [][]int{{1}},
		},
		{
			list: []interface{}{},
			res:  [][]int{},
		},
	}

	for _, data := range datas {
		node := genTreeNode(data.list)
		res := LevelOrder(node)
		assert.Equal(t, data.res, res)
		res2 := LevelOrder(node)
		assert.Equal(t, data.res, res2)
	}

}
