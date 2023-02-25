package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MaxInt, math.MinInt)
}

func isValidBSTHelper(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTHelper(root.Left, root.Val, min) && isValidBSTHelper(root.Right, max, root.Val)
}
