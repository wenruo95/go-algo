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

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNodes(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i := 0; i < len(vals); i++ {
		nodes[i] = &TreeNode{Val: vals[i]}
	}

	for i := 0; i < len(vals); i++ {
		if nodes[i] == nil {
			continue
		}

		if leftIndex := 2*i + 1; leftIndex < len(vals) && nodes[leftIndex] != nil {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex := 2*i + 2; rightIndex < len(vals) && nodes[rightIndex] != nil {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	return nodes[0]
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, root.Val)
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)
	return result
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)
	return result
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, PostorderTraversal(root.Left)...)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// leetcode 95: https://leetcode.com/problems/unique-binary-search-trees-ii/
func GenerateTrees(n int) []*TreeNode {
	set := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		set[i] = struct{}{}
	}

	var fn func() [][]int
	fn = func() [][]int {
		if len(set) <= 1 {
			var val int
			for item := range set {
				val = item
			}
			return [][]int{{val}}
		}

		result := make([][]int, 0)
		for val := 1; val <= n; val++ {
			if _, exist := set[val]; !exist {
				continue
			}

			delete(set, val)
			for _, v2 := range fn() {
				list := append([]int{val}, v2...)
				result = append(result, list)
			}
			set[val] = struct{}{}
		}

		return result
	}

	keys := make(map[string]struct{})
	nodes := make([]*TreeNode, 0)
	for _, v2 := range fn() {
		node := GenerateBinarySearchTree(v2...)

		key := fmt.Sprintf("%+v %+v %+v", PreorderTraversal(node), InorderTraversal(node), PostorderTraversal(node))
		if _, exist := keys[key]; exist {
			continue
		}
		keys[key] = struct{}{}
		nodes = append(nodes, node)
	}
	return nodes
}

func GenerateBinarySearchTree(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	for i := 1; i < len(vals); i++ {
		node := root
		for {

			if vals[i] > node.Val {
				if node.Right == nil {
					node.Right = &TreeNode{Val: vals[i]}
					break
				}
				node = node.Right
				continue
			}

			if node.Left == nil {
				node.Left = &TreeNode{Val: vals[i]}
				break
			}
			node = node.Left
		}

	}

	return root
}
