package leetcode

func numTrees(n int) int {
	buf := make([]int, n+1)
	buf[0] = 1
	var getBuf func(int) int
	getBuf = func(n int) int {
		if buf[n] == 0 {
			result := 0
			for i := 1; i <= n; i++ {
				l := getBuf(i - 1)
				r := getBuf(n - i)
				result += l * r
			}
			buf[n] = result
		}
		return buf[n]
	}
	return getBuf(n)
}
