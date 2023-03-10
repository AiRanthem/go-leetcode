package leetcode

/*
首先不考虑冻结期，对于某一天结束后，股票只有两种情况，持有或未持有，则需要分两种状态进行DP。
持有：当天买入或是之前买入后今天空过
未持有：当天卖掉或是之前卖掉之后今天空过

hold[i] = max(hold[i-1], sold[i-1]-prices[i])
sold[i] = max(hold[i-1] + price[i], sold[i-1])

考虑冻结期，未持有发生了变化：
未持有：前一天卖掉了或是今天卖掉了或是空过
分歧点在于前一天卖掉还是今天卖掉。尝试写sold的转移方程：

sold[i] = max(hold[i-1] + price[i], sold[i-1])

发现空过的选项（sold[i-1]）没有办法反应是不是前一天卖掉的
因此需要细化状态。
持有：今天不是冷冻期且今天购买；之前持有今天空过
未持有：之前持有今天卖出；今天冷冻期只能空过；今天正常但是选择空过
从“持有”的状态中我们发现了第三种状态，就是“今天不是冷冻期”，也就是“昨天没有卖出”。必须要定位到这个状态才能写出转移方程。
因此我们要调整状态sold，缩小为今天卖出，因为如果处于这种状态，则第二天无法买入。
则剩余情况为昨天卖了导致今天冷冻期只能空过或是今天正常但是选择空过，总之命名为未持有空过（skip）。那么整理状态转移方程如下：

sold[i] = hold[i-1] + price[i] // 缩小范围后的sold
hold[i] = max(hold[i-1], skip[i-1]-prices[i])
skip[i] = max(sold[i-1], skip[i-1])
*/
// func maxProfit(prices []int) int {
// 	hold, sold, skip := make([]int, len(prices)), make([]int, len(prices)), make([]int, len(prices))
// 	hold[0] = -prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		hold[i] = maxInt(hold[i-1], skip[i-1]-prices[i]) // 持有状态，之前已经买入或者今天买入
// 		sold[i] = hold[i-1] + prices[i]                  // 因卖出而未持有
// 		skip[i] = maxInt(skip[i-1], sold[i-1])           // 未持有且没有操作，说明前一天卖出或也未操作1
// 	}
// 	return maxInt(maxInt(hold[len(prices)-1], sold[len(prices)-1]), skip[len(prices)-1])
// }
