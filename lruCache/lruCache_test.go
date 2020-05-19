package lruCache

import "testing"

func TestLRUCache(t *testing.T) {
	res := make([]int, 0)
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	res = append(res, cache.Get(1))
	cache.Put(3, 3)
	res = append(res, cache.Get(2))
	cache.Put(4, 4)
	res = append(res, cache.Get(1))
	res = append(res, cache.Get(3))
	res = append(res, cache.Get(4))

	data := []int{1,-1,-1,3,4}
	flag := true

	for i := range res {
		if res[i] != data[i] {
			flag = false
			break
		}
	}

	if !flag {
		t.Errorf("Failed")
	} else {
		t.Logf("success")
	}
}