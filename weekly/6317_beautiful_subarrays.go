package weekly

func beautifulSubarrays(nums []int) (cnt int64) {
	finder := make(map[int][]int)
	p := 0
	for i := 0; i < len(nums); i++ {
		p ^= nums[i]
		if p == 0 {
			cnt++
		}
		cnt += int64(len(finder[p]))
		finder[p] = append(finder[p], i)
	}
	return
}
