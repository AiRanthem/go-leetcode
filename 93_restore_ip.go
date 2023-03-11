package leetcode

import "strings"

func restoreIpAddresses(s string) (result []string) {
	var path []string
	var dfs func(start, n int)
	dfs = func(start, n int) {
		if n == 3 {
			last := atoi(s[start:])
			if last >= 0 && last <= 255 && !(len(s)-start > 1 && s[start] == '0') {
				path = append(path, s[start:])
				result = append(result, strings.Join(path, "."))
				path = path[:len(path)-1]
			}
			return
		}
		// x.xxxxx
		if len(s)-start > 1 {
			path = append(path, s[start:start+1])
			dfs(start+1, n+1)
			path = path[:len(path)-1]
		}
		// xx.xxxx
		if len(s)-start > 2 && s[start] != '0' {
			path = append(path, s[start:start+2])
			dfs(start+2, n+1)
			path = path[:len(path)-1]
		}
		// xxx
		if len(s)-start > 3 && s[start] != '0' && atoi(s[start:start+3]) <= 255 {
			path = append(path, s[start:start+3])
			dfs(start+3, n+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}
