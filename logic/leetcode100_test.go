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
		res2 := LevelOrder2(node)
		assert.Equal(t, data.res, res2)
	}

}

func TestZigzagLevelOrder(t *testing.T) {

	type testData struct {
		list []interface{}
		res  [][]int
	}

	datas := []*testData{
		{
			list: []interface{}{3, 9, 20, nil, nil, 15, 7},
			res:  [][]int{{3}, {20, 9}, {15, 7}},
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
		res := ZigzagLevelOrder(node)
		assert.Equal(t, data.res, res)
	}

}

func TestMaxDepth(t *testing.T) {

	type testData struct {
		list  []interface{}
		depth int
	}

	datas := []*testData{
		{
			list:  []interface{}{3, 9, 20, nil, nil, 15, 7},
			depth: 3,
		},
		{
			list:  []interface{}{1, nil, 2},
			depth: 2,
		},
	}

	for _, data := range datas {
		node := genTreeNode(data.list)
		depth := MaxDepth(node)
		assert.Equal(t, data.depth, depth)
	}

}

func TestPreOrderAndInOrderBuildTree(t *testing.T) {
	type testData struct {
		preorder []int
		inorder  []int
		result   []interface{}
	}

	datas := []*testData{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			result:   []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			result:   []interface{}{-1},
		},
	}

	for _, data := range datas {
		tree := PreOrderAndInOrderBuildTree(data.preorder, data.inorder)
		node := genTreeNode(data.result)
		assert.Equal(t, node, tree)
		assert.True(t, IsSameTree(tree, node))

	}
}

func TestInOrderAndPostOrderBuildTree(t *testing.T) {
	type testData struct {
		inorder   []int
		postorder []int
		result    []interface{}
	}

	datas := []*testData{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			result:    []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			inorder:   []int{-1},
			postorder: []int{-1},
			result:    []interface{}{-1},
		},
	}

	for _, data := range datas {
		tree := InOrderAndPostOrderBuildTree(data.inorder, data.postorder)
		node := genTreeNode(data.result)
		assert.Equal(t, node, tree)
		assert.True(t, IsSameTree(tree, node))

	}
}

func TestLevelOrderBottom(t *testing.T) {

	type testData struct {
		list []interface{}
		res  [][]int
	}

	datas := []*testData{
		{
			list: []interface{}{3, 9, 20, nil, nil, 15, 7},
			res:  [][]int{{15, 7}, {9, 20}, {3}},
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
		res := LevelOrderBottom(node)
		assert.Equal(t, data.res, res)
	}

}
