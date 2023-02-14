package leetcode

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 4, 2}, []int{1, 2, 4}},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums1: []int{1, 3, 7, 1, 7, 5},
				nums2: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
