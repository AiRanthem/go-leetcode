package leetcode

// func maxProfit(prices []int) int {
// 	sold, hold := make([]int, len(prices)), make([]int, len(prices))
// 	hold[0] = -prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		hold[i] = maxInt(hold[i-1], sold[i-1]-prices[i])
// 		sold[i] = maxInt(sold[i-1], hold[i-1]+prices[i])
// 	}
// 	return maxInt(sold[len(prices)-1], hold[len(prices)-1])
// }
