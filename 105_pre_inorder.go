package leetcode

// pre:	[root, [pre of left], [pre of right]]
// in:	[[in of left], root, [in of right]]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	i := 0
	for inorder[i] != preorder[0] {
		i++
	}
	// now, "i" is the idx of root at inorder.
	node := TreeNode{Val: preorder[0]}
	node.Left = buildTree(preorder[1:i+1], inorder[0:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return &node
}
