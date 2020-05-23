package frequencySort

import "container/heap"

func frequencySort(s string) string {
	a := []byte(s)
	myMap := make(map[byte]int)
	for _, v := range a {
		if value, ok := myMap[v]; ok {
			myMap[v] = value + 1
		} else {
			myMap[v] = 1
		}
	}
	q := make(PriorityQueue, len(myMap))
	i := 0
	for value, priority := range myMap {
		q[i] = &Item{
			value: string(value),
			priority: priority,
			index: i,
		}
		i++
	}
	heap.Init(&q)
	var res string
	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		f := item.priority
		for i := 1; i <= f; i++ {
			res += item.value
		}
	}
	return res
}

type Item struct {
	value string
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}