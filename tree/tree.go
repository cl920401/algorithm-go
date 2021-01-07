package tree

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type MyNode struct {
	*TreeNode
	Level int
}

func levelOrderBottom(root *TreeNode) [][]int {
	nodes := make([][]int, 0)
	floorQueue := list.New()

	floorQueue.PushBack(&MyNode{
		root, 0,
	})

	for floorQueue.Len() != 0 {
		node := floorQueue.Remove(floorQueue.Front()).(*MyNode)
		if node != nil && node.TreeNode != nil {
			floorQueue.PushBack(&MyNode{
				node.Left, node.Level + 1,
			})
			floorQueue.PushBack(&MyNode{
				node.Right, node.Level + 1,
			})

			if len(nodes) <= node.Level {
				nodes = append(nodes, []int{node.Val})
			} else {
				nodes[node.Level] = append(nodes[node.Level], node.Val)
			}
		}
	}
	return nodes
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
