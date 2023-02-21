package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test",
			args{[]int{1, 3}, []int{2}},
			2.0,
		},
		{
			"test",
			args{[]int{1, 2}, []int{3, 4}},
			2.5,
		},
		{
			"test",
			args{[]int{1}, []int{2, 3, 4, 5, 6}},
			3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
