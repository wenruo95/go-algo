package logic

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// leetcode 90: https://leetcode.com/problems/subsets-ii/
func SubsetsWithDup(nums []int) [][]int {

	var (
		track  = make([]int, len(nums))
		result = make([][]int, 0)
	)
	sort.Ints(nums)

	var backtrack func(start int, tlen int)
	backtrack = func(start int, high int) {
		result = append(result, append([]int{}, track[0:high]...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			track[high] = nums[i]
			backtrack(i+1, high+1)
		}
	}

	backtrack(0, 0)

	return result
}

// leetcode 91: https://leetcode.com/problems/decode-ways/
func NumDecodings(s string) int {

	var (
		dp   func(s string, start int) int
		memo = make(map[int]int)
	)

	dp = func(s string, start int) int {
		if start > len(s)-1 {
			return 1
		}
		if s[start] <= '0' || s[start] > '9' {
			return 0
		}
		if v, exist := memo[start]; exist {
			return v
		}

		count := dp(s, start+1)
		if start+1 < len(s) {
			if sum := (s[start]-'0')*10 + (s[start+1] - '0'); sum > 0 && sum <= 26 {
				count = count + dp(s, start+2)
			}
		}

		memo[start] = count
		return count
	}

	return dp(s, 0)
}

// leetcode 93: https://leetcode.com/problems/restore-ip-addresses/
func RestoreIpAddresses(s string) []string {
	var (
		pos [3]int
		dp  func(s string, index, posIndex int) []string
	)

	dp = func(s string, index int, posIndex int) []string {
		if posIndex >= 3 || index >= len(s) {
			if size := len(s) - index; (size == 1) ||
				(size == 2 && s[index] != '0') || (size == 3 && s[index] != '0' && s[index:] <= "255") {
				return []string{s[0:pos[0]] + "." + s[pos[0]:pos[1]] + "." + s[pos[1]:pos[2]] + "." + s[pos[2]:]}
			}
			return nil
		}

		pos[posIndex] = index + 1
		arr := dp(s, index+1, posIndex+1)

		if s[index] != '0' {
			if index+1 < len(s) {
				pos[posIndex] = index + 2
				arr = append(arr, dp(s, index+2, posIndex+1)...)
			}

			if index+2 < len(s) && s[index:index+3] <= "255" {
				pos[posIndex] = index + 3
				arr = append(arr, dp(s, index+3, posIndex+1)...)
			}
		}

		return arr
	}

	return dp(s, 0, 0)
}

// leetcode 95: https://leetcode.com/problems/unique-binary-search-trees-ii/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {

	var generate func(lo, hi int) []*TreeNode

	generate = func(lo, hi int) []*TreeNode {
		if lo > hi {
			return []*TreeNode{nil}
		}
		if lo == hi {
			return []*TreeNode{{Val: lo}}
		}

		res := make([]*TreeNode, 0)
		for i := lo; i <= hi; i++ {
			leftList := generate(lo, i-1)
			rightList := generate(i+1, hi)
			for _, left := range leftList {
				for _, right := range rightList {
					root := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}
					res = append(res, root)
				}
			}

		}
		return res
	}

	return generate(1, n)
}

// leetcode 96: https://leetcode.com/problems/unique-binary-search-trees/description/
func NumTrees(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[i] + dp[j]*dp[i-j-1]
		}
	}

	return dp[n]
}

// leetcode 97: https://leetcode.com/problems/interleaving-string/
func IsInterleave(s1 string, s2 string, s3 string) bool {

	var judge func(l1, l2, l3 int) bool
	memo := make(map[string]bool)

	judge = func(l1, l2, l3 int) bool {
		if l1 == len(s1) && l2 == len(s2) && l3 == len(s3) {
			return true
		}
		if len(s3) == l3 || len(s3)-l3 != len(s2)-l2+len(s1)-l1 {
			return false
		}

		key := strconv.Itoa(l1) + "_" + strconv.Itoa(l2) + "_" + strconv.Itoa(l3)
		if v, exist := memo[key]; exist {
			return v
		}

		result := (l1 < len(s1) && s1[l1] == s3[l3] && judge(l1+1, l2, l3+1)) ||
			(l2 < len(s2) && s2[l2] == s3[l3] && judge(l1, l2+1, l3+1))
		memo[key] = result

		return result
	}

	return judge(0, 0, 0)
}

// leetcode 98: https://leetcode.com/problems/validate-binary-search-tree/
func IsValidBST(root *TreeNode) bool {
	var checkFn func(node *TreeNode, min, max int) bool
	checkFn = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}

		return checkFn(node.Left, min, intMin(max, node.Val)) &&
			checkFn(node.Right, intMax(min, node.Val), max)
	}

	return checkFn(root, math.MinInt, math.MaxInt)
}

// leetcode 99: https://leetcode.com/problems/recover-binary-search-tree/
func RecoverTree(root *TreeNode) {
	var recov func(node *TreeNode) (*TreeNode, *TreeNode)

	recov = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}

		leftMin, leftMax := recov(node.Left)
		rightMin, rightMax := recov(node.Right)
		if leftMax != nil && node.Val < leftMax.Val {
			fmt.Printf("[SWAP] node:%v leftmax:%v\n", node.Val, leftMax.Val)
			node.Val, leftMax.Val = leftMax.Val, node.Val
		}
		if rightMin != nil && node.Val > rightMin.Val {
			fmt.Printf("[SWAP] node:%v rightmin:%v\n", node.Val, rightMin.Val)
			node.Val, rightMin.Val = rightMin.Val, node.Val
		}

		if leftMin == nil || leftMin.Val > node.Val {
			leftMin = node
		}
		if rightMax == nil || rightMax.Val < node.Val {
			rightMax = node
		}
		return leftMin, rightMax
	}

	recov(root)
}
