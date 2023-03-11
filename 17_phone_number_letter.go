package leetcode

import "strings"

func letterCombinations(digits string) (result []string) {
	if len(digits) == 0 {
		return
	}
	var phoneNumberMap = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var path []string
	var dfs func(n int)
	dfs = func(n int) {
		if n == len(digits) {
			result = append(result, strings.Join(path, ""))
		} else {
			todo := phoneNumberMap[digits[n]]
			for i := 0; i < len(todo); i++ {
				path = append(path, todo[i])
				dfs(n + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
