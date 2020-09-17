package linklist

import (
	"testing"
)

func TestCombineLink(t *testing.T) {
	type args struct {
		root1 *ListNode
		root2 *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCombineLink",
			args: args{
				root1: &ListNode{
					next: &ListNode{
						next: &ListNode{
							next: nil,
							val:  6,
						},
						val: 3,
					},
					val: 1,
				},
				root2: &ListNode{
					next: &ListNode{
						next: &ListNode{
							next: nil,
							val:  5,
						},
						val: 4,
					},
					val: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombineLink(tt.args.root1, tt.args.root2)

			t.Log(got)
		})
	}
}
