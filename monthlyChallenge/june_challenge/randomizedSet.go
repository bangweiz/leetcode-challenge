package june_challenge

import "math/rand"

type RandomizedSet struct {
	set map[int]int
	arr []int
}


func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}


func (r *RandomizedSet) Insert(val int) bool {
	_, ok := r.set[val]
	if ok {
		return false
	}
	r.set[val] = len(r.set)
	r.arr = append(r.arr, val)
	return true
}


func (r *RandomizedSet) Remove(val int) bool {
	v, ok := r.set[val]
	if ok {
		r.set[r.arr[len(r.arr) - 1]] = v
		r.arr[v] = r.arr[len(r.arr) - 1]
		r.arr = r.arr[:len(r.arr) - 1]
		delete(r.set, val)
		return true
	}
	return false
}


func (r *RandomizedSet) GetRandom() int {
	return r.arr[rand.Intn(len(r.arr))]
}
