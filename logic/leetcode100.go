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
