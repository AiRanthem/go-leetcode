package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		return (findTopKSmall(nums1, nums2, l/2) + findTopKSmall(nums1, nums2, l/2+1)) / 2
	} else {
		return findTopKSmall(nums1, nums2, l/2+1)
	}
}

func findTopKSmall(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if k == 1 {
		if len(nums2) > 0 && nums1[0] > nums2[0] {
			return float64(nums2[0])
		} else {
			return float64(nums1[0])
		}
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	p := k / 2
	var i, j int
	if len(nums2) < p {
		j = len(nums2) - 1
		i = k - j - 2
	} else {
		i = p - 1
		j = p - 1
	}
	if nums1[i] < nums2[j] {
		return findTopKSmall(nums1[i+1:], nums2, k-i-1)
	} else {
		return findTopKSmall(nums1, nums2[j+1:], k-j-1)
	}
}
