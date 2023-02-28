package leetcode

func sortColors(nums []int) {
	p0, p1 := 0, 0 // px: 第一个可能不为x的数
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 { // 如果p0<p1，那么一定p0指向1，所以要把换出去的1还给p1
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if num == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
