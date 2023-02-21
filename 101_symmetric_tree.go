package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return compareSynnetric(root.Left, root.Right)
}

func compareSynnetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	outside := compareSynnetric(left.Left, right.Right)
	inside := compareSynnetric(left.Right, right.Left)
	return outside && inside && left.Val == right.Val
}
