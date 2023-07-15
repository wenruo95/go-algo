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
			list:  []interface{}{2, nil, 3, nil, 4, nil, 6, nil, 6},
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

func TestHasPathSum(t *testing.T) {

	type testData struct {
		list   []interface{}
		target int
		has    bool
	}

	datas := []*testData{
		{
			list:   []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1},
			target: 22,
			has:    true,
		},
		{
			list:   []interface{}{1, 2, 3},
			target: 5,
			has:    false,
		},
		{
			list:   []interface{}{},
			target: 0,
			has:    false,
		},
		{
			list:   []interface{}{1, 2},
			target: 1,
			has:    false,
		},
	}
	for _, data := range datas {
		node := genTreeNode(data.list)
		has := HasPathSum(node, data.target)
		assert.Equal(t, data.has, has)
	}

}

func TestPathSum(t *testing.T) {

	type testData struct {
		list   []interface{}
		target int
		res    [][]int
	}

	datas := []*testData{
		{
			list:   []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
			target: 22,
			res:    [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			list:   []interface{}{1, 2, 3},
			target: 5,
			res:    [][]int{},
		},
		{
			list:   []interface{}{1, 2},
			target: 0,
			res:    [][]int{},
		},
	}
	for _, data := range datas {
		node := genTreeNode(data.list)
		res := PathSum(node, data.target)
		assert.Equal(t, data.res, res)
	}

}

func TestFlattern(t *testing.T) {
	type testData struct {
		list []interface{}
		res  []interface{}
	}

	datas := []*testData{
		{
			list: []interface{}{1, 2, 5, 3, 4, nil, 6},
			res:  []interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6},
		},
		{
			list: []interface{}{},
			res:  []interface{}{},
		},

		{
			list: []interface{}{1},
			res:  []interface{}{1},
		},
	}

	for _, data := range datas {
		root := genTreeNode(data.list)
		Flatten(root)
		rootList := dumpTreeNode(root)
		assert.Equal(t, data.res, rootList)
	}

}

func TestNumDistinct(t *testing.T) {
	type testData struct {
		s, t string
		res  int
	}

	datas := []*testData{
		{s: "rabbbit", t: "rabbit", res: 3},
		{s: "babgbag", t: "bag", res: 5},
	}
	for _, data := range datas {
		res := NumDistinct(data.s, data.t)
		assert.Equal(t, data.res, res)
	}
}
