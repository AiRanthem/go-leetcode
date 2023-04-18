package main

import (
	"fmt"
	"math"
	"strings"
)

// 删除字符串中出现次数最少的字符
func main() {
	var s string
	_, _ = fmt.Scan(&s)
	counter := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	var mn = math.MaxInt
	for _, n := range counter {
		if n < mn {
			mn = n
		}
	}
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if counter[s[i]] != mn {
			sb.WriteByte(s[i])
		}
	}
	fmt.Println(sb.String())
}
