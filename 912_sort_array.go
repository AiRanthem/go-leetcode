package leetcode

import "math/rand"

func sortArray(nums []int) []int {
	iSort(nums)
	return nums
}

func iSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		pre := i - 1
		cur := nums[i]
		for pre >= 0 && nums[pre] > cur {
			nums[pre+1] = nums[pre]
			pre--
		}
		nums[pre+1] = cur
	}
}

func qSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	l, r := left, right
	p := int(rand.Uint32())%(right-left+1) + left
	nums[left], nums[p] = nums[p], nums[left]
	pix := nums[left]
	for l < r {
		for l < r && nums[r] > pix {
			r--
		}
		for l < r && nums[l] <= pix {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[left] = nums[left], nums[l]
	qSort(nums, left, l-1)
	qSort(nums, l+1, right)
}
