package leetcode

/*
多状态DP
1. 持有，第一次持有：今天买；之前买
2. 不持有，第一次卖出：今天卖，之前卖
3. 持有，第二次持有：今天买；之前买
4. 不持有，第二次卖出：今天卖，之前卖
（第一次持有前收益恒为0，不考虑）
*/

func maxProfit(prices []int) int {
	n := len(prices)
	hold1 := make([]int, n)
	hold2 := make([]int, n)
	sold1 := make([]int, n)
	sold2 := make([]int, n)
	// sold0 理解为全都是0就行了，省略
	hold1[0] = -prices[0]
	hold2[0] = -prices[0] // 第一天买入，卖掉，再买入(-_-)
	for i := 1; i < n; i++ {
		hold1[i] = maxInt(hold1[i-1], -prices[i]) // sold0[i] = 0
		sold1[i] = maxInt(sold1[i-1], hold1[i-1]+prices[i])
		hold2[i] = maxInt(hold2[i-1], sold1[i-1]-prices[i])
		sold2[i] = maxInt(sold2[i-1], hold2[i-1]+prices[i])
	}
	return maxInt(sold1[n-1], sold2[n-1])
}
