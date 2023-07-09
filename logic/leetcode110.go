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
