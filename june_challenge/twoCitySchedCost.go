package june_challenge

import "container/heap"

func twoCitySchedCost(costs [][]int) int {
	q := make(PriorityQueue, len(costs))
	for i := range costs {
		q[i] = &Item{i, costs[i][1] - costs[i][0], i}
	}
	heap.Init(&q)
	res := 0
	for len(q) >  len(costs) / 2 {
		item := heap.Pop(&q).(*Item)
		res += costs[item.vaule][0]
	}
	for len(q) > 0 {
		item := heap.Pop(&q).(*Item)
		res += costs[item.vaule][1]
	}
	return res
}

type Item struct {
	vaule int
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