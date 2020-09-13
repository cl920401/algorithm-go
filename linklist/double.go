package linklist

type DoubleLinkList struct {
	head *DoubleListNode
	tail *DoubleListNode
	len  int
}

type DoubleListNode struct {
	prev *DoubleListNode
	next *DoubleListNode
	val  interface{}
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (v *DoubleLinkList) searchNode(n int) *DoubleListNode {
	if n >= v.len {
		return nil
	}
	search := v.head
	for i := 0; i < n; i++ {
		search = search.next
	}
	return search
}

func (v *DoubleLinkList) Insert(n int, val interface{}) {
	if n == 0 {
		node := DoubleListNode{
			next: v.head,
			val:  val,
		}
		v.head = &node
		v.tail = &node
		v.len++
		return
	}
	search := v.searchNode(n - 1)
	if search == nil {
		return
	}
	node := DoubleListNode{
		next: search.next,
		prev: search,
		val:  val,
	}
	if search.next != nil {
		search.next.prev = &node
	} else {
		v.tail = &node
	}
	search.next = &node
	v.len++
}

func (v *DoubleLinkList) Delete(n int) {
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
	if search.next.next != nil {
		search.next.next.prev = search
	} else {
		v.tail = search
	}
	v.len--
}

func (v *DoubleLinkList) GetLength() int {
	return v.len
}

func (v *DoubleLinkList) Search(val interface{}) int {
	search := v.head
	for i := 0; i < v.len; i++ {
		if val == search.val {
			return i
		}
		search = search.next
	}
	return -1
}

func (v *DoubleLinkList) isNull() bool {
	return v.len == 0
}
