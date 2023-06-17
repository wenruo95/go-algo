package logic

// leetcode 100: https://leetcode.com/problems/same-tree/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// leetcode 101: https://leetcode.com/problems/symmetric-tree/
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	return check(root.Left, root.Right)
}

// leetcode 102: https://leetcode.com/problems/binary-tree-level-order-traversal/
func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		list2 := make([]*TreeNode, 0)
		valList := make([]int, 0)

		var valid bool
		for _, v := range list {
			if v == nil {
				continue
			}
			valList = append(valList, v.Val)
			valid = valid || v.Left != nil || v.Right != nil
			list2 = append(list2, v.Left, v.Right)
		}

		if len(valList) > 0 {
			res = append(res, valList)
		}
		if !valid {
			break
		}
		list = list2[0:]
	}

	return res
}

func LevelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := NewQueue()
	queue.Push(root)
	for queue.Top() != nil {
		size := queue.Size()
		list := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.Pop().(*TreeNode)
			list = append(list, node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		res = append(res, list)
	}

	return res
}

// leetcode 103: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := NewQueue()
	queue.Push(root)
	for queue.Top() != nil {
		size := queue.Size()
		list := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.Pop().(*TreeNode)
			list = append(list, node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}

		if len(res)%2 != 0 {
			for i := 0; i < len(list)/2; i++ {
				j := len(list) - 1 - i
				list[i], list[j] = list[j], list[i]
			}
		}
		res = append(res, list)
	}

	return res
}

// leetcode 104: https://leetcode.com/problems/maximum-depth-of-binary-tree/
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return intMax(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

// leetcode 105: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func PreOrderAndInOrderBuildTree(preorder []int, inorder []int) *TreeNode {
	var build func(preLeft, preRight, inLeft, inRight int) *TreeNode

	inValue2Index := make(map[int]int)
	for index, value := range inorder {
		inValue2Index[value] = index
	}

	build = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight || inLeft > inRight {
			return nil
		}
		node := new(TreeNode)
		node.Val = preorder[preLeft]

		pivot := inValue2Index[node.Val]
		leftCount := pivot - inLeft

		node.Left = build(preLeft+1, preLeft+leftCount, inLeft, pivot-1)
		node.Right = build(preLeft+leftCount+1, preRight, pivot+1, inRight)
		return node
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

// leetcode 106: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func InOrderAndPostOrderBuildTree(inorder []int, postorder []int) *TreeNode {
	var build func(inLeft, inRight, postLeft, postRight int) *TreeNode

	inValue2Index := make(map[int]int)
	for index, value := range inorder {
		inValue2Index[value] = index
	}

	build = func(inLeft, inRight, postLeft, postRight int) *TreeNode {
		if inLeft > inRight || postLeft > postRight {
			return nil
		}
		node := new(TreeNode)
		node.Val = postorder[postRight]

		pivot := inValue2Index[node.Val]
		leftCount := pivot - inLeft

		node.Left = build(inLeft, pivot-1, postLeft, postLeft+leftCount-1)
		node.Right = build(pivot+1, inRight, postLeft+leftCount, postRight-1)
		return node
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
}
