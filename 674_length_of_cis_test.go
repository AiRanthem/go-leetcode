package leetcode

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{[]int{1, 3, 5, 4, 7}},
			want: 3,
		},
		{
			name: "test",
			args: args{[]int{2, 2, 2, 2, 2}},
			want: 1,
		},
		{
			name: "test",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
