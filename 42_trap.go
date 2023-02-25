package leetcode

func trap(height []int) int {
	var result int
	l, r := 0, len(height)-1
	lh, rh := height[l], height[r]
	for l < r {
		for l < r && lh <= rh {
			l++
			if height[l] > lh {
				lh = height[l]
			} else {
				result += lh - height[l]
			}
		}
		for l < r && rh <= lh {
			r--
			if height[r] > rh {
				rh = height[r]
			} else {
				result += rh - height[r]
			}
		}
	}
	return result
}
