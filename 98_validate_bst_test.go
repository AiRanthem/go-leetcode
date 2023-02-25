package leetcode

import "testing"

func Test98(t *testing.T) {
	isValidBST(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}})
}
