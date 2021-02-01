package array

import (
	"container/heap"
)

type bigHeap []int

func (b bigHeap) Len() int {
	return len(b)
}

func (b bigHeap) Less(i, j int) bool {
	return b[i] > b[j]
}

func (b bigHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *bigHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *bigHeap) Pop() interface{} {
	if b.Len() == 0 {
		return 0
	}
	defer func() {
		*b = (*b)[:len(*b)-1]
	}()
	return (*b)[len(*b)-1]
}

func (b *bigHeap) Peek() int {
	if b.Len() == 0 {
		return 0
	}
	return (*b)[0]
}

type littleHeap []int

func (l littleHeap) Len() int {
	return len(l)
}

func (l littleHeap) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l littleHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *littleHeap) Push(x interface{}) {
	*l = append(*l, x.(int))
}

func (l *littleHeap) Pop() interface{} {
	if l.Len() == 0 {
		return 0
	}
	defer func() {
		*l = (*l)[:len(*l)-1]
	}()
	return (*l)[len(*l)-1]
}

func (l *littleHeap) Peek() int {
	if l.Len() == 0 {
		return 0
	}
	return (*l)[0]
}

type MedianFinder struct {
	left  bigHeap
	right littleHeap
}

/** initialize your data structure here. */
func ConstructorMedianFinder() MedianFinder {
	mid := MedianFinder{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
	heap.Init(&mid.left)
	heap.Init(&mid.right)
	return mid
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 && this.right.Len() == 0 {
		heap.Push(&this.left, num)
		return
	}
	if float64(num) > this.FindMedian() {
		heap.Push(&this.right, num)
	} else {
		heap.Push(&this.left, num)
	}
	this.Rebalance()
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(this.left.Peek()+this.right.Peek()) / 2
	} else if this.left.Len() > this.right.Len() {
		return float64(this.left.Peek())
	} else {
		return float64(this.right.Peek())
	}
}

func (this *MedianFinder) Rebalance() {
	for this.left.Len() > 0 && this.left.Len() > this.right.Len()+1 {
		heap.Push(&this.right, heap.Pop(&this.left))
	}
	for this.right.Len() > 0 && this.left.Len()+1 < this.right.Len() {
		heap.Push(&this.left, heap.Pop(&this.right))
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
