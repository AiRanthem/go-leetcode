package leetcode

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	waiting := 0
	maxIncome, rollTimes, curIncome := 0, 0, 0
	for i := 0; i < len(customers) || waiting > 0; i++ {
		if i < len(customers) {
			waiting += customers[i]
		}
		curIncome -= runningCost
		if waiting > 4 {
			curIncome += 4 * boardingCost
			waiting -= 4
		} else {
			curIncome += waiting * boardingCost
			waiting = 0
		}
		if curIncome > maxIncome {
			maxIncome = curIncome
			rollTimes = i + 1
		}
	}
	if maxIncome <= 0 {
		return -1
	}
	return rollTimes
}
