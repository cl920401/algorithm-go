package linklist

import "container/list"

type MinStack struct {
	MinList   *list.List
	StackList *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		MinList:   list.New(),
		StackList: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	if this.StackList.Len() == 0 || x <= this.Min() {
		this.MinList.PushBack(x)
	}
	this.StackList.PushBack(x)

}

// 协程不安全
func (this *MinStack) Pop() {
	if this.StackList.Len() == 0 {
		return
	}
	e := this.StackList.Remove(this.StackList.Back()).(int)
	if e == this.Min() {
		this.MinList.Remove(this.MinList.Back())
	}
}

func (this *MinStack) Top() int {
	if this.Len() == 0 {
		return 0
	}
	return this.StackList.Back().Value.(int)
}

func (this *MinStack) Min() int {
	if this.Len() == 0 {
		return 0
	}
	return this.MinList.Back().Value.(int)
}

func (this *MinStack) Len() int {
	return this.StackList.Len()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
