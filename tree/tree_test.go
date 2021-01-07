package tree

import (
	"container/list"
	"reflect"
	"testing"
)

func GetTreeNode(arr []int) *TreeNode {
	queue := list.New()
	root := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	queue.PushBack(root)
	for i := 1; i < len(arr); i+=2 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if arr[i] >= 0 {
			node.Left = &TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			queue.PushBack(node.Left)
		}
		if arr[i+1] >= 0 {
			node.Right = &TreeNode{
				Val:   arr[i+1],
				Left:  nil,
				Right: nil,
			}
			queue.PushBack(node.Right)
		}
	}
	return root
}


func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_levelOrderBottom",
			args: args{
				root: GetTreeNode([]int{41,37,44,24,39,42,48,1,35,38,40,-1,43,46,49,0,2,30,36,-1,-1,-1,-1,-1,-1,45,47,-1,-1,-1,-1,-1,4,29,32,-1,-1,-1,-1,-1,-1,3,9,26,-1,31,34,-1,-1,7,11,25,27,-1,-1,33,-1,6,8,10,16,-1,-1,-1,28,-1,-1,5,-1,-1,-1,-1,-1,15,19,-1,-1,-1,-1,12,-1,18,20,-1,13,17,-1,-1,22,-1,14,-1,-1,21,23}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(perOrder(tt.args.root))
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}

		})
	}
}
