package leetcode

type TreeQueue2 []*TreeNode

func (q *TreeQueue2) Len() int {
	return len(*q)
}

func (q *TreeQueue2) PushRight(node *TreeNode) {
	*q = append(*q, node)
}

func (q *TreeQueue2) PushLeft(node *TreeNode) {
	*q = append([]*TreeNode{node}, *q...)
}

func (q *TreeQueue2) PopLeft() *TreeNode {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *TreeQueue2) PopRight() *TreeNode {
	val := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return val
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	q := TreeQueue2{}
	if root != nil {
		q.PushRight(root)
	}
	goRight := true
	var result [][]int
	for q.Len() != 0 {
		n := q.Len()
		var row []int
		for i := 0; i < n; i++ {
			if goRight {
				node := q.PopLeft()
				row = append(row, node.Val)
				if node.Left != nil {
					q.PushRight(node.Left)
				}
				if node.Right != nil {
					q.PushRight(node.Right)
				}
			} else {
				node := q.PopRight()
				row = append(row, node.Val)
				if node.Right != nil {
					q.PushLeft(node.Right)
				}
				if node.Left != nil {
					q.PushLeft(node.Left)
				}
			}
		}
		goRight = !goRight
		result = append(result, row)
	}
	return result
}
