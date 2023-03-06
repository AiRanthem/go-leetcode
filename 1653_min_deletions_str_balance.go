package leetcode

// dp[i]: 前i位
// if s[i] == 'b': dp[i] = 	dp[i-1]
// if s[i] == 'a': dp[i] = 	min{ dp[i-1]+1 , numB (累计B的数量）}
//
// 从左往右遍历，可以降维优化
func minimumDeletions1(s string) int {
	numB := 0
	dp := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			numB++
		} else {
			dp = minInt(dp+1, numB)
		}
	}
	return dp
}

// 前后缀分解
// 平衡串的前缀：全是a，后缀：全是b
// 操作：去掉前缀所有的b，后缀所有的a，遍历所有分割点。
// 第一次遍历：分割点从0开始，全是后缀，统计后面的a的个数，记为当前删除数
// 第二次遍历：分割点开始后移，如果遇到a，则当前删除数量-1（后缀少删除一个a，前缀不用动）
// 如果遇到b，则当前删除量+1（前缀多删一个b，后缀不用动）
func minimumDeletions2(s string) int {
	toDel := 0
	for i := 0; i < len(s); i++ {
		toDel += int('b' - s[i]) // count 'a'
	}
	ans := toDel
	for i := 0; i < len(s); i++ {
		toDel += int(s[i]-'a')*2 - 1 // 'a' == -1, 'b' == 1
		if toDel < ans {
			ans = toDel
		}
	}

	return ans
}

func minimumDeletions(s string) int {
	return minimumDeletions2(s)
}
