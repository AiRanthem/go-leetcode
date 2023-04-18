package classic150

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	i, j := m-1, m+n-1
	for i >= 0 {
		nums1[j] = nums1[i]
		j--
		i--
	}
	i, j, k := n, 0, 0
	for i < m+n && j < n {
		if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}
