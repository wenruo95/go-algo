package logic

import (
	"strconv"
)

// leetcode 110: https://leetcode.com/problems/balanced-binary-tree/
func IsBalanced(root *TreeNode) bool {
	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)
		if left == -1 || right == -1 || right-left > 1 || left-right > 1 {
			return -1
		}
		return intMax(left, right) + 1
	}

	return height(root) != -1
}

// leetcode 111: https://leetcode.com/problems/minimum-depth-of-binary-tree/
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MinDepth(root.Left)
	right := MinDepth(root.Right)

	if left == 0 || right == 0 {
		return intMax(left, right) + 1
	}
	return intMin(left, right) + 1
}

// leetcode 112: https://leetcode.com/problems/path-sum/
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && (root.Left == nil && root.Right == nil) {
		return true
	}
	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}

// leetcode 113: https://leetcode.com/problems/path-sum-ii/
func PathSum(root *TreeNode, targetSum int) [][]int {

	var findPath func(node *TreeNode, target int) [][]int
	findPath = func(node *TreeNode, target int) [][]int {
		if node == nil {
			return [][]int{}
		}
		if node.Val == target && node.Left == nil && node.Right == nil {
			return [][]int{{node.Val}}
		}
		lres := findPath(node.Left, target-node.Val)
		rres := findPath(node.Right, target-node.Val)

		res := make([][]int, 0)
		for i := 0; i < len(lres); i++ {
			res = append(res, append([]int{node.Val}, lres[i]...))
		}
		for i := 0; i < len(rres); i++ {
			res = append(res, append([]int{node.Val}, rres[i]...))
		}

		return res
	}
	return findPath(root, targetSum)
}

// leetcode 114: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
func Flatten(root *TreeNode) {
	validNode := func(nodes ...*TreeNode) *TreeNode {
		for _, node := range nodes {
			if node != nil {
				return node
			}
		}
		return nil
	}
	var reorder func(node *TreeNode) (*TreeNode, *TreeNode)

	reorder = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}

		lfirst, llast := reorder(node.Left)
		rfirst, rlast := reorder(node.Right)
		if node.Left != nil {
			node.Left = nil
			node.Right = lfirst
		}
		llast = validNode(llast, node)
		if node.Right != nil {
			llast.Left = nil
			llast.Right = rfirst
		}

		return node, validNode(rlast, llast, node)
	}

	reorder(root)
}

// leetcode 115: https://leetcode.com/problems/distinct-subsequences/
func NumDistinct(s string, t string) int {

	memo := make(map[string]int)

	var dp func(si, ti int) int
	dp = func(si, ti int) int {
		if ti == len(t) {
			return 1
		}
		if si >= len(s) {
			return 0
		}

		key := strconv.Itoa(si) + "_" + strconv.Itoa(ti)
		if value, exist := memo[key]; exist {
			return value
		}

		var res int
		if s[si] == t[ti] {
			res = dp(si+1, ti+1) + dp(si+1, ti)
		} else {
			res = dp(si+1, ti)
		}
		memo[key] = res
		return res
	}

	return dp(0, 0)
}

type B116Node struct {
	Val   int
	Left  *B116Node
	Right *B116Node
	Next  *B116Node
}

// leetcode 116: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
func Connect(root *B116Node) *B116Node {
	if root == nil {
		return root
	}
	queue := NewQueue()
	queue.Push(root)
	for queue.Size() > 0 {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node := queue.Pop().(*B116Node)
			if i != size-1 {
				node.Next = queue.Top().(*B116Node)
			}
			if node.Left == nil || node.Right == nil {
				continue
			}
			queue.Push(node.Left)
			queue.Push(node.Right)
		}
	}
	return root
}
