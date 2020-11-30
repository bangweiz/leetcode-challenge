package contest199

import (
	"container/heap"
)

type Item struct {
	value byte
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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

func (pq *PriorityQueue) update(item *Item, value byte, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func restoreString(s string, indices []int) string {
	b, pq := []byte(s), make(PriorityQueue, len(indices))
	for i, v := range indices {
		pq[i] = &Item{b[i], v, i}
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		b[item.priority] = item.value
	}
	return string(b)
}

func minFlips(target string) int {
	b, prev, res := []byte(target), byte('0'), 0
	for _, v := range b {
		if v != prev {
			res, prev = res + 1, v
		}
	}
	return res
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	res := new(int)
	dfs(root, distance, res)
	return *res
}

func dfs(root *TreeNode, distance int, count *int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{1}
	}
	left := dfs(root.Left, distance, count)
	right := dfs(root.Right, distance, count)

	for i := range left {
		for j := range right {
			if left[i] + right[j] < distance {
				*count++
			}
		}
	}

	left = append(left, right...)
	for i, v := range left {
		left[i] = v + 1
	}
	return left
}

func getLengthOfOptimalCompression(s string, k int) int {
	memo := make([][]int, 101)
	for i := range memo {
		memo[i] = make([]int, 101)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return recursion(s, 0, k, memo)
}

func recursion(s string, start, k int, memo [][]int) int {
	if k < 0 {
		return 999
	}
	if len(s) - start <= k {
		return 0
	}
	if memo[start][k] != -1 {
		return memo[start][k]
	}

	count, remain := 0, k
	res := recursion(s, start + 1, remain - 1, memo)

	for i := start; i < len(s); i++ {
		if s[i] == s[start] {
			count++
		} else {
			remain--
			if remain < 0 {
				break
			}
		}
		res = min(res, recursion(s, i + 1, remain, memo) + calLength(count))
	}
	memo[start][k] = res
	return res
}

func calLength(count int) int {
	if count > 100 {
		return 4
	} else if count > 10 {
		return 3
	} else if count > 1 {
		return 2
	}
	return 1
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}