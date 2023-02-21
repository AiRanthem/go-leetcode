package leetcode

type TreeQueue []*TreeNode

func (q *TreeQueue) Push(x *TreeNode) {
	*q = append(*q, x)
}

func (q *TreeQueue) Pop() *TreeNode {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *TreeQueue) Len() int {
	return len(*q)
}

func levelOrder(root *TreeNode) [][]int {
	queue := &TreeQueue{}
	if root != nil {
		queue.Push(root)
	}
	var result [][]int
	for queue.Len() > 0 {
		var row []int
		l := queue.Len()
		for i := 0; i < l; i++ {
			n := queue.Pop()
			row = append(row, n.Val)
			if n.Left != nil {
				queue.Push(n.Left)
			}
			if n.Right != nil {
				queue.Push(n.Right)
			}
		}
		result = append(result, row)
	}
	return result
}
