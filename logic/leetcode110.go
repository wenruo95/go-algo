package logic

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
