package leetcode

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		traversal(root, &result)
	}
	return result
}

func traversal(node *TreeNode, result *[]int) {
	if node.Left != nil {
		traversal(node.Left, result)
	}
	if node.Right != nil {
		traversal(node.Right, result)
	}
	*result = append(*result, node.Val)
}
