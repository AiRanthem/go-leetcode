package leetcode

//func isMatch(s string, p string) bool {
//	pl, sl := len(p), len(s)
//	if pl == 0 {
//		return sl == 0
//	}
//	if pl > 1 && p[1] == '*' {
//		k := 0
//		for k < sl && matchOne(s[k], p[0]) {
//			k++
//		}
//		for i := 0; i <= k; i++ {
//			if ok := isMatch(s[i:], p[2:]); ok {
//				return true
//			}
//		}
//		return false
//	} else if sl > 0 && matchOne(s[0], p[0]) {
//		return isMatch(s[1:], p[1:])
//	} else {
//		return false
//	}
//}
//
//func matchOne(a, b byte) bool {
//	return a == b || a == '.' || b == '.'
//}

/*
dp[i][j]; s[:i] matches p[j]
if p[j] is letter:

	dp[i][j] = dp[i-1][j-1]	; s[i] "matches" p[j]
	dp[i][j] = false 		; else

if p[j] is star:

	dp[i][j] = dp[i-1][j]（匹配1个后保持j继续匹配） || dp[i][j-2]（匹配0个，相当于这个*不存在）	; s[i] "matches" p[j-1]
	dp[i][j] = dp[i][j-2]（只能匹配0个）	; else
*/
func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)
	dp := make([][]bool, sl+1)
	for i := 0; i <= sl; i++ {
		dp[i] = make([]bool, pl+1)
	}
	dp[0][0] = true
	for i := 0; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if p[j-1] != '*' {
				if i > 0 && regexMatch(s[i-1], p[j-1]) {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			} else {
				if i > 0 && regexMatch(s[i-1], p[j-2]) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[sl][pl]
}

func regexMatch(a, b byte) bool {
	return a == b || b == '.'
}
