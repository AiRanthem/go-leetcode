package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}
