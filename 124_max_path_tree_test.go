package leetcode

import "testing"

func Test124(t *testing.T) {
	maxPathSum(&TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	})
}
