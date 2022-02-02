/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : tree.go
*   coder: zemanzeng
*   date : 2022-01-21 16:54:53
*   desc : 二叉树
*
================================================================*/

package tree

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func New(val interface{}) *TreeNode {
	if val == nil {
		return nil
	}

	node := new(TreeNode)
	node.Val = val
	return node
}

func NewTreeNodes(values ...interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(values))
	for i := 0; i < len(values); i++ {
		nodes[i] = New(values[i])
	}

	for i := 0; i < len(values); i++ {
		if nodes[i] == nil {
			continue
		}

		if leftIndex := 2*i + 1; leftIndex < len(values) && nodes[leftIndex] != nil {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex := 2*i + 2; rightIndex < len(values) && nodes[rightIndex] != nil {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	return nodes[0]
}

func PreorderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	result := make([]interface{}, 0)
	result = append(result, root.Val)
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)
	return result
}

func InorderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	result := make([]interface{}, 0)
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)
	return result
}

func PostorderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	result := make([]interface{}, 0)
	result = append(result, PostorderTraversal(root.Left)...)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
