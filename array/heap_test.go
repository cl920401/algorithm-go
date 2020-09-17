package array

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want KthLargest
	}{
		{
			name: "TestConstructor",
			args: args{
				k:    2,
				nums: []int{0},
			},
			want: KthLargest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor(tt.args.k, tt.args.nums)
			t.Log(*got.littleHeap)
			got.Add(-1)
			t.Log(*got.littleHeap)
			got.Add(1)
			t.Log(*got.littleHeap)
			got.Add(-2)
		})
	}
}
