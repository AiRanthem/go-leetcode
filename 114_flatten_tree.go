package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var helper func(node *TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		left, right := node.Left, node.Right
		var lastLeft, lastRight *TreeNode
		if left == nil && right == nil {
			return node
		}
		if left != nil {
			lastLeft = helper(left)
			lastLeft.Right = right
			node.Right = left
			node.Left = nil
		}
		if right != nil {
			lastRight = helper(right)
			return lastRight
		} else {
			return lastLeft
		}
	}
	helper(root)
}
