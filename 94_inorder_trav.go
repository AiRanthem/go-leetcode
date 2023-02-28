package leetcode

var result94 []int

func inorderTraversal(root *TreeNode) []int {
	result94 = []int{}
	inorderTraversalHelper(root)
	return result94
}

func inorderTraversalHelper(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalHelper(root.Left)
	result94 = append(result94, root.Val)
	inorderTraversalHelper(root.Right)
}
