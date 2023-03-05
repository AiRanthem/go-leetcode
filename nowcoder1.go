package leetcode

func sol1(building []int, D int) int {
	N := len(building)

	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N && building[j]-building[i] <= D; j++ {
			for k := j + 1; k < N && building[k]-building[j] <= D; k++ {
				ans++
			}
		}
	}
	return ans
}
