package logic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestRestoreIpAddresses(t *testing.T) {
	type testData struct {
		s      string
		result []string
	}
	datas := []*testData{
		{
			s:      "25525511135",
			result: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:      "0000",
			result: []string{"0.0.0.0"},
		},
		{
			s:      "101023",
			result: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			s:      "1111",
			result: []string{"1.1.1.1"},
		},
	}

	for _, data := range datas {
		if result := RestoreIpAddresses(data.s); !stringListItemEqual(data.result, result) {
			t.Errorf("restore_ip_address error. s:%v result:%v expect:%v", data.s, result, data.result)
		}
	}

}

func TestNumTrees(t *testing.T) {

	type testData struct {
		n      int
		result int
	}

	datas := []*testData{
		{n: 0, result: 1},
		{n: 1, result: 1},
		{n: 2, result: 2},
		{n: 3, result: 5},
	}

	for _, data := range datas {
		result := NumTrees(data.n)
		assert.Equal(t, data.result, result)
	}

}

func TestIsInterleave(t *testing.T) {
	type testData struct {
		s1, s2, s3 string
		expect     bool
	}

	datas := []*testData{
		{
			s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac",
			expect: true,
		},
		{
			s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc",
			expect: false,
		},
		{
			s1: "", s2: "", s3: "",
			expect: true,
		},
		{
			s1:     "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
			s2:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			s3:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			expect: false,
		},
		{
			s1:     "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
			s2:     "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
			s3:     "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			expect: false,
		},
	}

	for _, data := range datas {
		result := IsInterleave(data.s1, data.s2, data.s3)
		assert.Equal(t, data.expect, result)
	}

}
func genTreeNode(list []interface{}) *TreeNode {
	if len(list) == 0 || list[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if v == nil {
			continue
		}

		node := &TreeNode{Val: v.(int)}
		if li := 2*i + 1; li < len(list) {
			node.Left = nodes[li]
		}
		if ri := 2*i + 2; ri < len(list) {
			node.Right = nodes[ri]
		}
		nodes[i] = node
	}
	fmt.Printf("[GEN] %+v\n", nodes)

	return nodes[0]
}

func TestIsValidBST(t *testing.T) {

	type testData struct {
		list  []interface{}
		valid bool
	}

	datas := []*testData{
		{
			list:  []interface{}{2, 1, 3},
			valid: true,
		},
		{
			list:  []interface{}{5, 1, 4, nil, nil, 3, 6},
			valid: false,
		},
		{
			list:  []interface{}{2, 2, 2},
			valid: false,
		},
	}
	for _, data := range datas {
		root := genTreeNode(data.list)
		result := IsValidBST(root)
		assert.Equal(t, data.valid, result)
	}
}

func treeNodeDump(root *TreeNode) []interface{} {
	nodes := make([]*TreeNode, 0)
	if root == nil {
		return []interface{}{}
	}
	nodes = append(nodes, root)

	list := []*TreeNode{root}
	list2 := make([]*TreeNode, 0)
	for len(list) > 0 {
		var hasValid bool
		for _, item := range list {
			if item == nil {
				list2 = append(list2, nil, nil)
				continue
			}

			hasValid = hasValid || item.Left != nil || item.Right != nil
			list2 = append(list2, item.Left, item.Right)
		}
		if !hasValid {
			break
		}

		nodes = append(nodes, list2...)
		list = list2[0:]
		list2 = make([]*TreeNode, 0)
	}
	fmt.Printf("[DEBUG-XXX] %+v [4]:%v\n", nodes, nodes[4])

	res := make([]interface{}, len(nodes))
	for i, v := range nodes {
		if v == nil {
			res[i] = nil
			continue
		}
		res[i] = v.Val
	}

	var pos int
	for pos = len(res) - 1; pos >= 0 && res[pos] == nil; pos-- {

	}

	fmt.Printf("[DEBUG-YYY] %+v %+v\n", res, res[0:pos+1])
	return res[0 : pos+1]
}

func TestRecoverTree(t *testing.T) {

	type testData struct {
		list   []interface{}
		result []interface{}
	}

	datas := []*testData{
		{
			list:   []interface{}{1, 3, nil, nil, 2},
			result: []interface{}{3, 1, nil, nil, 2},
		},
		{
			list:   []interface{}{3, 1, 4, nil, nil, 2},
			result: []interface{}{2, 1, 4, nil, nil, 3},
		},
	}

	for _, data := range datas {
		root := genTreeNode(data.list)
		t.Logf("[DEBUG0] %+v", treeNodeDump(root))
		RecoverTree(root)
		t.Logf("[DEBUG1] %+v", treeNodeDump(root))

		result := treeNodeDump(root)
		assert.Equal(t, data.result, result)
	}
}
