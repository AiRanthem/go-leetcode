package leetcode

import "testing"

func Test_lastStoneWeightII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{[]int{2, 7, 4, 1, 8, 1}},
			1,
		},
		{
			"test",
			args{[]int{31, 26, 33, 21, 40}},
			5,
		},
		{
			"test",
			args{[]int{1, 2}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightII(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeightII() = %v, want %v", got, tt.want)
			}
		})
	}
}
