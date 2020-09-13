package LRU

import "container/list"

type LRUCache struct {
	max     int
	cache   map[int]*list.Element
	lruList *list.List
}

type kv struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		max:     capacity,
		cache:   make(map[int]*list.Element),
		lruList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		if realValue, ok := elem.Value.(*kv); ok {
			this.lruList.MoveToFront(elem)
			return realValue.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	newValue := kv{
		key:   key,
		value: value,
	}
	if elem, ok := this.cache[key]; ok {
		this.lruList.MoveToFront(elem)
		elem.Value = &newValue
	} else {
		this.cache[key] = this.lruList.PushFront(&newValue)
		if this.lruList.Len() > this.max {
			this.RemoveOldest()
		}
	}
}

func (this *LRUCache) RemoveOldest() {
	if this.cache == nil {
		return
	}
	if ele := this.lruList.Back(); ele != nil {
		this.lruList.Remove(ele)
		if kv, ok := ele.Value.(*kv); ok {
			delete(this.cache, kv.key)
		}
	}
}
