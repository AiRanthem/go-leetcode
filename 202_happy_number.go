package leetcode

func isHappy(n int) bool {
	set := make(map[int]struct{})
	for true {
		if n == 1 {
			break
		}
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = struct{}{}
		s := 0
		for n > 0 {
			i := n % 10
			s += i * i
			n /= 10
		}
		n = s
	}
	return true
}
