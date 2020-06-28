package contest195

import (
	"container/heap"
	"math"
	"strconv"
)

func isPathCrossing(path string) bool {
	b, hmap, start := []byte(path), map[string]int{"0,0": 1}, []int{0, 0}
	for _, v := range b {
		if v == 'N' {
			start[0] += 1
		} else if v == 'S' {
			start[0] -= 1
		} else if v == 'W' {
			start[1] -= 1
		} else {
			start[1] += 1
		}
		temp := strconv.Itoa(start[0]) + "," + strconv.Itoa(start[1])
		if _, ok := hmap[temp]; ok {
			return true
		} else {
			hmap[temp] = 1
		}
	}
	return false
}

func canArrange(arr []int, k int) bool {
	hmap := make(map[int]int, 0)
	for i := 0; i < k; i++ {
		hmap[i] = 0
	}

	for _, v := range arr {
		temp := v % k
		if temp < 0 {
			temp += k
		}
		hmap[temp] = hmap[temp] + 1
	}

	if hmap[0] % 2 != 0 {
		return false
	}

	if k % 2 == 0 && hmap[k / 2] % 2 != 0 {
		return false
	}

	for i := 1; i <= k / 2; i++ {
		if hmap[i] != hmap[k - i] {
			return false
		}
	}
	return true
}

func numSubseq(nums []int, target int) int {
	nums = mergeSort(nums)
	count, res := -2, 0
	for i := range nums {
		if nums[0] + nums[i] > target {
			count = i - 1
			break
		}
	}
	if count == -1 {
		 return 0
	}
	if count == -2 {
		count = len(nums) - 1
	}
	power := []int{1}
	for i := 1; i <= count; i++ {
		power = append(power, power[i - 1] * 2 % 1000000007)
	}
	for i := 0; i <= count; i++ {
		for nums[i] + nums[count - i] > target  {
			count--
		}
		if count - i >= 0 {
			res = (res + power[count]) % 1000000007
		}
	}
	return res
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(a, b []int) []int {
	c, i, j, index := make([]int, len(a) + len(b)), 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[index] = a[i]
			i++
		} else {
			c[index] = b[j]
			j++
		}
		index++
	}

	for i < len(a) {
		c[index] = a[i]
		i++
		index++
	}

	for j < len(b) {
		c[index] = b[j]
		j++
		index++
	}
	return c
}

type Item struct {
	value int
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func findMaxValueOfEquation(points [][]int, k int) int {
	res := math.MinInt64
	qp := make(PriorityQueue, 0)
	heap.Init(&qp)
	for i, point := range points {
		if len(qp) == 0 {
			heap.Push(&qp, &Item{point[0], point[1] - point[0], i})
			continue
		}
		for len(qp) > 0 && point[0] - qp[len(qp) - 1].value > k {
			heap.Pop(&qp)
		}
		if len(qp) > 0 {
			res = max(res, qp[len(qp) - 1].priority + point[0] + point[1])
		}
		heap.Push(&qp, &Item{point[0], point[1] - point[0], i})
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
