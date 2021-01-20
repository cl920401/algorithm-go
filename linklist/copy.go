package linklist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mapNode := make(map[*Node]*Node)
	next := head
	for next != nil {
		mapNode[next] = &Node{
			Val:    next.Val,
			Next:   next.Next,
			Random: next.Random,
		}
		next = next.Next
	}
	next = head
	for next != nil {
		mapNode[next].Next = mapNode[next.Next]
		mapNode[next].Random = mapNode[next.Random]
		next = next.Next
	}
	return mapNode[head]
}

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return root.Left
	}
	return root
}
