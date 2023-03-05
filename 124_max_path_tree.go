package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var through func(node *TreeNode) int
	through = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val
		leftContribution := through(node.Left)
		rightContribution := through(node.Right)
		if leftContribution > 0 {
			sum += leftContribution
		}
		if rightContribution > 0 {
			sum += rightContribution
		}
		if sum > result {
			result = sum
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
	through(root)
	return result
}
