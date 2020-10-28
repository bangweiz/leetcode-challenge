package contest198

import (
	"math"
	"sort"
)

func numWaterBottles(numBottles, numExchange int) (res int) {
	res = numBottles
	for numBottles >= numExchange {
		a, b := numBottles / numExchange, numBottles % numExchange
		res, numBottles = res + a, a + b
	}
	return
}

func countSubTrees(n int, edges [][]int, labels string) (res []int) {
	tree, b, hmap, graph, isVisited := make([][]int, n), []byte(labels), make([]map[byte]int, n), make([][]int, n), make([]bool, n)
	for i := range tree {
		tree[i], graph[i], hmap[i] = make([]int, 0), make([]int, 0), map[byte]int{b[i]: 1}
	}
	for _, v := range edges {
		tree[v[0]] = append(tree[v[0]], v[1])
		tree[v[1]] = append(tree[v[1]], v[0])

	}

	buildEdges(tree, graph, isVisited, 0)
	dfs(graph, 0, b, hmap);

	for i, v := range hmap {
		res = append(res, v[b[i]])
	}

	return
}

func buildEdges(tree, graph [][]int, isVisited [] bool, current int) {
	isVisited[current] = true
	for _, v := range tree[current] {
		if !isVisited[v] {
			graph[current] = append(graph[current], v)
			buildEdges(tree, graph, isVisited, v)
		}
	}
}

func dfs(tree [][]int, current int, labels []byte, hmap []map[byte]int) {
	for _, v := range tree[current] {
		dfs(tree, v, labels, hmap)
	}
	for _, v := range tree[current] {
		for key, value := range hmap[v] {
			if count, ok := hmap[current][key]; ok {
				hmap[current][key] = count + value
			} else {
				hmap[current][key] = value
			}
		}
	}
}

type Item struct {
	left, right int
}

type SubStrings []Item

func (s SubStrings) Len() int {
	return len(s)
}

func (s SubStrings) Less(i, j int) bool {
	return s[i].right - s[i].left + 1 < s[j].right - s[j].left + 1
}

func (s SubStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func maxNumOfSubstrings(s string) (res []string) {
	b, left, right, substr := []byte(s), make(map[byte]int), make(map[byte]int), make(SubStrings, 0)
	for i, v := range b {
		if _, ok := left[v]; !ok {
			left[v] = i
		}
		right[v] = i
	}

	for key, l := range left {
		flag := true
		r, _ := right[key]
		for i := l; i < r; i++ {
			if index, _ := left[b[i]]; index < l {
				flag = false
				break
			}
			if index, _ := right[b[i]]; index > r {
				r = index
			}
		}
		if flag {
			substr = append(substr, Item{l, r})
		}
	}

	sort.Sort(substr)
	resItems := make(SubStrings, 0)
	for _, s := range substr {
		flag := true
		for _, item := range resItems {
			if !(item.right < s.left || item.left > s.right) {
				flag = false
				break
			}
		}
		if flag {
			resItems = append(resItems, s)
		}
	}

	for _, s := range resItems {
		res = append(res, string(b[s.left:s.right + 1]))
	}
	return
}

type Node struct {
	Low, High, Value int
	Left, Right *Node
}

func (n *Node) query(low, high int) int {
	if n.Low == low && n.High == high {
		return n.Value
	}
	mid := (n.Low + n.High) / 2
	if high <= mid {
		return n.Left.query(low, high)
	}
	if low > mid {
		return n.Right.query(low, high)
	}
	return n.Left.query(low, mid) & n.Right.query(mid + 1, high)
}

func closestToTarget(arr []int, target int) int {
	tree := buildSegmentTree(arr, 0, len(arr) - 1)
	res := math.MaxInt64
	for l, _ := range arr {
		r := biSearch(tree, l, len(arr) - 1, target)
		res = int(math.Min(float64(res) ,math.Abs(float64(target - tree.query(l, r)))))
		if r < len(arr) - 1 {
			res = int(math.Min(float64(res), float64(target - tree.query(l, r + 1))))
		}
		if res == 0 {
			return 0
		}
	}
	return res
}

func biSearch(root *Node, start, end, target int) int {
	lo, hi := start, end
	if root.query(lo, lo) <= target {
		return lo
	}
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if root.query(start, mid) >= target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func buildSegmentTree(arr []int, low, high int) *Node {
	if low == high {
		return &Node{low, high, arr[low], nil, nil}
	}
	mid := (low + high) / 2
	left, right := buildSegmentTree(arr, low, mid), buildSegmentTree(arr, mid + 1, high)
	n := &Node{low, high, left.Value & right.Value, left, right}
	return n
}
