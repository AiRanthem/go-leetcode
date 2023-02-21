package leetcode

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := j * minInt(height[i], height[j])
	for i < j {
		hi, hj := height[i], height[j]
		if height[i] < height[j] {
			for i < j && height[i] <= hi {
				i++
			}
		} else {
			for i < j && height[j] <= hj {
				j--
			}
		}
		area = maxInt(area, (j-i)*minInt(height[i], height[j]))
	}
	return area
}
