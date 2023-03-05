package leetcode

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	ans := 0
	for n, _ := range set {
		l := 1
		for i, ok := 1, true; ok; i++ {
			_, ok = set[n+i]
			if ok {
				delete(set, n+i)
				l++
			}
		}
		for i, ok := 1, true; ok; i++ {
			_, ok = set[n-i]
			if ok {
				delete(set, n+i)
				l++
			}
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
