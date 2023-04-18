package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		m[s[i]] = struct{}{}
	}
	fmt.Println(len(m))
}
