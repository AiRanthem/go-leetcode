package leetcode

import "strings"

func printBin(num float64) string {
	sb := strings.Builder{}
	sb.WriteByte('0')
	sb.WriteByte('.')
	for i := 0; i <= 30; i++ {
		num *= 2
		if num >= 1 {
			num -= 1
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
		if num == 0 {
			return sb.String()
		}
	}
	return "ERROR"
}
