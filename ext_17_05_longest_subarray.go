package leetcode

func findLongestSubarray(array []string) []string {
	var maxL, maxR, sum int
	maxR = -1
	firstOccur := make(map[int]int)
	firstOccur[0] = -1
	for i := 0; i < len(array); i++ {
		if array[i][0] < '=' {
			sum--
		} else {
			sum++
		}
		if l, ok := firstOccur[sum]; ok {
			if maxR-maxL+1 < i-l || (maxR-maxL+1 == i-l && l+1 < maxL) {
				maxR = i
				maxL = l + 1
			}
		} else {
			firstOccur[sum] = i
		}
	}
	return array[maxL : maxR+1]
}
