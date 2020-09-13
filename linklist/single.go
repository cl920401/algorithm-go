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
