package daily

func jump(nums []int) (steps int) {
	l, r, next := 0, 0, 0 // 这一步最远跳到r, 下一步最远跳到next
	for r < len(nums)-1 {
		for l <= r {
			if l+nums[l] > next {
				next = l + nums[l]
			}
			l++
		}
		steps++
		r = next
	}
	return
}
