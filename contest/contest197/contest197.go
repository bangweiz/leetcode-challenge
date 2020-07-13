package contest197

import (
	"container/heap"
	"math"
)

func numIdenticalPairs(nums []int) (res int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return
}

func numSub(s string) (res int) {
	b, count, temp, max := []byte(s), make([]int, 0), 0, 0
	for i := range b {
		if b[i] == '1' {
			temp++
		} else if temp != 0 {
			if max < temp {
				max = temp
			}
			count = append(count, temp)
			temp = 0
		}
	}
	if temp != 0 {
		if max < temp {
			max = temp
		}
		count = append(count, temp)
	}
	if len(count) != 0 {
		r := make([]int, max)
		r[0] = 1
		for i := 1; i < len(r); i++ {
			r[i] = (r[i - 1] + i + 1) % 1000000007
		}

		for _, v := range count {
			res += r[v - 1] % 1000000007
		}
	}
	return
}

type Item struct {
	value int
	priority float64
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

func (pq *PriorityQueue) update(item *Item, value int, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	res, q, graph, visited := make([]float64, n), make(PriorityQueue, 1), make([]map[int]float64, n), make([]bool, n)
	for i := range res {
		res[i] = 0
	}
	res[start] = 1
	for i := range graph {
		graph[i] = make(map[int]float64)
	}
	for i := range edges {
		graph[edges[i][0]][edges[i][1]] = succProb[i]
		graph[edges[i][1]][edges[i][0]] = succProb[i]
	}
	q[0] = &Item{start, 1, 0}
	heap.Init(&q)

	for q.Len() > 0 {
		current := heap.Pop(&q).(*Item)
		if !visited[current.value] {
			visited[current.value] = true
			for next, pos := range graph[current.value] {
				temp := current.priority * pos
				if temp > res[next] {
					res[next] = temp
					heap.Push(&q, &Item{next, temp, q.Len()})
				}
			}
		}
	}
	return res[end]
}


// TODO: not pass
func getMinDistSum(positions [][]int) (res float64) {
	if len(positions) == 2 {
		res = math.Sqrt(math.Pow(float64(positions[0][0] - positions[1][0]), 2) + math.Pow(float64(positions[0][1] - positions[1][1]), 2))
	} else if len(positions) > 2 {
		res = math.MaxFloat64
		x, y, delta, minX, minY := 50.0, 50.0, 50.0, 50.0, 50.0
		for delta >= 0.00001 {
			for i := x - delta; i <= x + delta; i += delta / 10 {
				for j := y - delta; j <= y + delta; j += delta / 10 {
					d := dist(positions, i, j)
					if d < res {
						res, minX, minY = d, i, j
					}
				}
			}
			x, y, delta = minX, minY, delta / 10
		}
	}
	return
}

func dist(positions [][]int, x, y float64) (d float64) {
	for _, v := range positions {
		d += math.Sqrt(math.Pow(float64(v[0]) - x, 2) + math.Pow(float64(v[1]) - y, 2))
	}
	return
}