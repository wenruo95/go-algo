/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : tree_test.go
*   coder: zemanzeng
*   date : 2022-02-02 10:26:36
*   desc :
*
================================================================*/

package tree

import "testing"

func TestGenerateTrees(t *testing.T) {
	m := map[int]int{
		1: 1,
		3: 5,
	}
	for n, length := range m {
		nodes := GenerateTrees(n)
		if len(nodes) != length {
			t.Errorf("generate_trees n:%v len:%v expected:%v", n, len(nodes), length)
			continue
		}
		for index, node := range nodes {
			t.Logf("k:%v index:%v pre_order:%+v", n, index, PreorderTraversal(node))
		}
	}
}
