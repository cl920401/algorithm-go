package search

import (
	"testing"
)

func TestCombine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "TestCombine",
			args: args{
				k: 2,
				n: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Combine(tt.args.k, tt.args.n))
		})
	}
}

func TestUpstairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upstairs(tt.args.n); got != tt.want {
				t.Errorf("Upstairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
