package top100

import "math"

func maxPathSum(root *TreeNode) (max int) {
	max = math.MinInt
	var contributionOf func(node *TreeNode) int
	contributionOf = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val
		leftContribution := contributionOf(node.Left)
		rightContribution := contributionOf(node.Right)
		if leftContribution > 0 {
			sum += leftContribution
		}
		if rightContribution > 0 {
			sum += rightContribution
		}
		if sum > max {
			max = sum
		}
		if leftContribution < rightContribution {
			leftContribution = rightContribution
		}
		if leftContribution > 0 {
			return node.Val + leftContribution
		} else {
			return node.Val
		}
	}
	_ = contributionOf(root)
	return
}
