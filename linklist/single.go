package linklist

type LinkList interface {
	Insert(i int, v interface{}) // 增
	Delete(i int)                // 删
	GetLength() int              // 获取长度
	Search(v interface{}) int    // 查
	isNull() bool                // 判断是否为空
}

type SingleLinkList struct {
	head *SingleListNode
	len  int
}

type SingleListNode struct {
	next *SingleListNode
	val  interface{}
}

func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{
		head: nil,
		len:  0,
	}
}

func (v *SingleLinkList) searchNode(n int) *SingleListNode {
	if n >= v.len {
		return nil
	}
	search := v.head
	for i := 0; i < n; i++ {
		search = search.next
	}
	return search
}

func (v *SingleLinkList) Insert(n int, val interface{}) {
	if n == 0 {
		node := SingleListNode{
			next: v.head,
			val:  val,
		}
		v.head = &node
		v.len++
		return
	}
	search := v.searchNode(n - 1)
	if search == nil {
		return
	}
	node := SingleListNode{
		next: search.next,
		val:  val,
	}
	search.next = &node
	v.len++
}

func (v *SingleLinkList) Delete(n int) {
	if n == 0 {
		v.head = nil
		v.len--
		return
	}
	search := v.searchNode(n - 1)
	if search == nil || search.next == nil {
		return
	}
	search.next = search.next.next
	v.len--
}

func (v *SingleLinkList) GetLength() int {
	return v.len
}

func (v *SingleLinkList) Search(val interface{}) int {
	search := v.head
	for i := 0; i < v.len; i++ {
		if val == search.val {
			return i
		}
		search = search.next
	}
	return -1
}

func (v *SingleLinkList) isNull() bool {
	return v.len == 0
}

type ListNode struct {
	next *ListNode
	rand *ListNode
	val  int
}

//CombineLink: 合并两个有序链表
func CombineLink(root1 *ListNode, root2 *ListNode) *ListNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	var root *ListNode
	var next1 *ListNode
	var next2 *ListNode
	if root2.val > root1.val {
		root = root1
		next1 = root1
		next2 = root2
	} else {
		root = root2
		next1 = root2
		next2 = root1
	}

	for next1.next != nil && next2 != nil {
		if next1.val > next2.val {
			return nil
		} else if next1.val <= next2.val && next1.next.val >= next2.val {
			next2, next1.next, next2.next = next2.next, next2, next1.next
		} else {
			next1 = next1.next
		}
	}
	if next1.next == nil {
		next1.next = next2
	}
	return root
}

//链表拷贝
func CopyLinkList(root *ListNode) *ListNode {
	next := root
	nodeMap := make(map[*ListNode]*ListNode)
	for next != nil {
		nodeMap[next] = &ListNode{
			next: next.next,
			rand: next.rand,
			val:  next.val,
		}
		next = next.next
	}
	next = root
	for next != nil {
		tmp := nodeMap[next]
		tmp.next = nodeMap[tmp.next]
		tmp.rand = nodeMap[tmp.rand]
		next = next.next
	}
	return nodeMap[root]
}
