package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			break
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
