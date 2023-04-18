package main

import (
	"fmt"
	"sort"
)

// 明明的随机数
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	m := make(map[int]struct{})
	var x int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)
		m[x] = struct{}{}
	}
	i := 0
	result := make([]int, len(m))
	for k, _ := range m {
		result[i] = k
		i++
	}
	sort.Ints(result)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
