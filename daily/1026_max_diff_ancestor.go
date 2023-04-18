package daily

import "math"

func maxAncestorDiff(root *TreeNode) (result int) {
	result = math.MinInt
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, max int, min int) {
		if node == nil {
			return
		}
		cur := maxInt(abs(node.Val-max), abs(node.Val-min))
		if cur > result {
			result = cur
		}
		if node.Val > max {
			max = node.Val
		}
		if node.Val < min {
			min = node.Val
		}
		dfs(node.Left, max, min)
		dfs(node.Right, max, min)
	}
	dfs(root, root.Val, root.Val)
	return
}
