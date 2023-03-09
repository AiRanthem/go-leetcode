package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]*ChildNode)
	for i := 0; i < len(prerequisites); i++ {
		a, b := get207Node(m, prerequisites[i][0]), get207Node(m, prerequisites[i][1])
		a.Children = append(a.Children, b)
	}
	var dfs func(depth int, node *ChildNode)
	fail := false
	dfs = func(depth int, node *ChildNode) {
		if depth > numCourses {
			fail = true
		}
		if node.Val == 1 {
			return
		}
		for i := 0; !fail && i < len(node.Children); i++ {
			dfs(depth+1, node.Children[i])
		}
		node.Val = 1
	}
	for i := 0; i < numCourses; i++ {
		if node, ok := m[i]; ok && node.Val == 0 {
			dfs(0, node)
			if fail {
				break
			}
		}
	}
	return !fail
}

func get207Node(m map[int]*ChildNode, n int) *ChildNode {
	if node, ok := m[n]; ok {
		return node
	} else {
		node = &ChildNode{}
		m[n] = node
		return node
	}
}
