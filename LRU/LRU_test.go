package LRU

import (
	"testing"
)

func TestLRU(t *testing.T) {
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_1 := obj.Get(1)
	t.Log(param_1)
	obj.Put(3, 3)
	param_1 = obj.Get(2)
	t.Log(param_1)
	obj.Put(4, 4)
	param_1 = obj.Get(1)
	t.Log(param_1)
	param_1 = obj.Get(3)
	t.Log(param_1)
	param_1 = obj.Get(4)
	t.Log(param_1)

}
