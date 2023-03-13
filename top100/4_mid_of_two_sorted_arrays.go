package top100

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if l := len(nums1) + len(nums2); l%2 == 1 {
		return findMedianSortedArraysKthSmall(nums1, nums2, l/2+1)
	} else {
		return (findMedianSortedArraysKthSmall(nums1, nums2, l/2+1) + findMedianSortedArraysKthSmall(nums1, nums2, l/2)) / 2
	}
}

func findMedianSortedArraysKthSmall(nums1 []int, nums2 []int, k int) float64 {
	reduce := k / 2
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums1) < reduce {
		reduce = len(nums1)
	}
	if reduce == 0 {
		return float64(minInt(nums1[0], nums2[0]))
	}
	if nums1[reduce-1] > nums2[reduce-1] {
		nums1, nums2 = nums2, nums1
	}
	return findMedianSortedArraysKthSmall(nums1[reduce:], nums2, k-reduce)
}
