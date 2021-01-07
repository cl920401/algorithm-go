package array

import "container/heap"

//设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
//你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。

type IntHeap []int

func (t IntHeap) Len() int {
	return len(t)
}

func (t IntHeap) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t IntHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (this *IntHeap) Push(n interface{}) {
	*this = append(*this, n.(int))
}

func (this *IntHeap) Pop() interface{} {
	old := *this
	*this = old[0 : len(*this)-1]
	return old[len(*this)-1]
}

type KthLargest struct {
	max        int
	littleHeap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	var littleHeap IntHeap
	if k > len(nums) {
		littleHeap = nums
	} else {
		littleHeap = nums[:k]
	}
	heap.Init(&littleHeap)
	kth := KthLargest{
		littleHeap: &littleHeap,
		max:        k,
	}
	if k < len(nums) {
		for _, num := range nums[k:] {
			kth.Add(num)
		}
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.littleHeap.Len() < this.max {
		heap.Push(this.littleHeap, val)
	} else if (*this.littleHeap)[0] < val {
		this.littleHeap.Push(val)
		heap.Remove(this.littleHeap, 0)
	}
	return (*this.littleHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
