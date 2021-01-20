package tree

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue []*TreeNode

func (t *queue) Push(node *TreeNode, isSingle bool) {
	if isSingle {
		if node.Left != nil {
			t.PushNode(node.Left)
		}
		if node.Right != nil {
			t.PushNode(node.Right)
		}
	} else {
		if node.Right != nil {
			t.PushNode(node.Right)
		}
		if node.Left != nil {
			t.PushNode(node.Left)
		}
	}
}

func (t *queue) PushNode(x *TreeNode) {
	*t = append(*t, x)
}

func (t *queue) Pop() *TreeNode {
	if len(*t) == 0 {
		return nil
	}
	res := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return res
}

func levelOrder(root *TreeNode) [][]int {
	i := make([][]int, 0)
	row := 0
	single := true
	var cur queue = make([]*TreeNode, 0)
	var next queue = make([]*TreeNode, 0)
	for root != nil {
		if len(i) == row {
			i = append(i, []int{})
		}
		i[row] = append(i[row], root.Val)
		next.Push(root, single)
		root = cur.Pop()
		if root == nil {
			row++
			single = !single
			root = next.Pop()
			cur, next = next, cur
		}

	}
	return i
}

func perOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := make([]int, 0)
	stack := list.New()
	for root != nil {
		if root.Left != nil {
			stack.PushBack(root)
			root = root.Left
			continue
		}
		nums = append(nums, root.Val)
		if root.Right != nil {
			root = root.Right
			continue
		}
		if stack.Len() <= 0 {
			break
		}
		for {
			root = stack.Remove(stack.Back()).(*TreeNode)
			if root == nil {
				break
			}
			nums = append(nums, root.Val)
			if root.Right != nil {
				root = root.Right
				break
			}
		}
	}
	return nums
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 || len(postorder) == 1 {
		return true
	}
	mid := postorder[len(postorder)-1]
	leftLen := -1
	for i := range postorder[:len(postorder)-1] {
		if postorder[i] > mid {
			leftLen = i
			break
		}
	}
	if leftLen < 0 {
		leftLen = len(postorder) - 1
	} else {
		for i := range postorder[leftLen+1 : len(postorder)-1] {
			if postorder[i+leftLen+1] < mid {
				return false
			}
		}
	}
	if !verifyPostorder(postorder[:leftLen]) {
		return false
	}

	if !verifyPostorder(postorder[leftLen : len(postorder)-1]) {
		return false
	}

	return true
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		if sum == 0 {
			return [][]int{{}}
		}
		return nil
	}
	pathLeft := pathSum(root.Left, sum-root.Val)
	if pathLeft != nil {
		for i := range pathLeft {
			pathLeft[i] = append([]int{root.Val}, pathLeft[i]...)
		}
	}
	pathRight := pathSum(root.Right, sum-root.Val)
	if pathRight != nil {
		for j := range pathRight {
			pathRight[j] = append([]int{root.Val}, pathRight[j]...)
		}
	}
	return append(pathLeft, pathRight...)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	val := strconv.Itoa(root.Val)
	nodes := []string{val}
	queueNode := list.New()
	queueNode.PushBack(root.Left)
	queueNode.PushBack(root.Right)

	for queueNode.Len() > 0 {
		root = queueNode.Remove(queueNode.Front()).(*TreeNode)
		if root == nil {
			val = "null"
			nodes = append(nodes, val)
		} else {
			val = strconv.Itoa(root.Val)
			nodes = append(nodes, val)
			queueNode.PushBack(root.Left)
			queueNode.PushBack(root.Right)
		}
	}
	return "[" + strings.Join(nodes, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) < 3 {
		return nil
	}
	nodes := strings.Split(data[1:len(data)-1], ",")
	if len(nodes) == 0 {
		return nil
	}
	queueNode := list.New()
	var root *TreeNode
	if i, err := strconv.Atoi(nodes[0]); err == nil {
		root = &TreeNode{
			Val:   i,
			Left:  nil,
			Right: nil,
		}
		queueNode.PushBack(root)
	} else {
		return nil
	}
	for i := 1; i < len(nodes); i += 2 {
		node := queueNode.Remove(queueNode.Front()).(*TreeNode)
		if val, err := strconv.Atoi(nodes[i]); err == nil {
			node.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			queueNode.PushBack(node.Left)
		}

		if val, err := strconv.Atoi(nodes[i+1]); err == nil {
			node.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			queueNode.PushBack(node.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
