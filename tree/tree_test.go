package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	code := Constructor()
	tree := code.deserialize("[1,2,3,4,5,6,7]")
	str := code.serialize(tree)
	fmt.Print(str)
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test_levelOrder",
			args: args{
				root: tree,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_verifyPostorder",
			args: args{
				postorder: []int{4, 8, 6, 12, 16, 14, 10},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
