package leetcode

type ParenthesisTree struct {
	Val   int // continous left
	Left  *ParenthesisTree
	Right *ParenthesisTree
}

func generateParenthesis(n int) []string {
	root := &ParenthesisTree{Val: 0}
	generateParentresisTree(root, n)
	var result []string
	buffer := make([]byte, 2*n)
	stack := make([]*ParenthesisTree, 2*n)
	stack[0] = root
	buffer[0] = '('
	cur := root.Left
	i := 1
	for cur != root {
		if i == 2*n {
			result = append(result, string(buffer))
		}
		if cur.Left != nil {
			buffer[i] = '('
			stack[i] = cur
			i++
			next := cur.Left
			cur.Left = nil
			cur = next
			continue
		}
		if cur.Right != nil {
			buffer[i] = ')'
			stack[i] = cur
			i++
			next := cur.Right
			cur.Right = nil
			cur = next
			continue
		}
		i--
		cur = stack[i]
	}
	return result
}

func generateParentresisTree(root *ParenthesisTree, n int) {
	if n > 0 {
		root.Left = &ParenthesisTree{Val: root.Val + 1}
		generateParentresisTree(root.Left, n-1)
	}
	if root.Val > 0 {
		root.Right = &ParenthesisTree{Val: root.Val - 1}
		generateParentresisTree(root.Right, n)
	}
}
