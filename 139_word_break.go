package leetcode

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	lens := make(map[int]struct{})

	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = struct{}{}
		lens[len(wordDict[i])] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for l, _ := range lens {
			if i >= l && dictContains(dict, s[i-l:i]) && dp[i-l] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func dictContains(dict map[string]struct{}, s string) bool {
	_, ok := dict[s]
	return ok
}
