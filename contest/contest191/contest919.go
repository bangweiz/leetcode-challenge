package contest191

import "math"

func maxProduct(nums []int) int {
	nums = mergeSort(nums)
	return (nums[1] - 1) * (nums[0] - 1)
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = mergeSort(horizontalCuts)
	verticalCuts = mergeSort(verticalCuts)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	max1, max2 := horizontalCuts[0], verticalCuts[0]
	for i := 1; i < len(horizontalCuts); i++ {
		if horizontalCuts[i] - horizontalCuts[i - 1] > max1  {
			max1 = horizontalCuts[i] - horizontalCuts[i - 1]
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		if verticalCuts[i] - verticalCuts[i - 1] > max2  {
			max2 = verticalCuts[i] - verticalCuts[i - 1]
		}
	}
	return max1 * max2 % int(math.Pow(10, 9) + 7)
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

func minReorder(n int, connections [][]int) int {
	road1, road2 := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		road1[i], road2[i] = make([]int, 0), make([]int, 0)
	}
	for _, v := range connections {
		road1[v[0]] = append(road1[v[0]], v[1])
		road2[v[1]] = append(road2[v[1]], v[0])
	}

	res := []int{0}
	q := []int{0}
	isVisited := make([]bool, n)
	isVisited[0] = true
	bfs(road1, road2, q, res, isVisited)
	return res[0]
}

func bfs(road1, road2 [][]int, q, res []int, isVisited []bool) {
	for len(q) > 0 {
		current := q[0]
		isVisited[current] = true
		q = q[1:]
		for i := 0; i < len(road1[current]); i++ {
			if !isVisited[road1[current][i]] {
				res[0] += 1
				q = append(q, road1[current][i])
			}
		}
		for i := 0; i < len(road2[current]); i++ {
			if !isVisited[road2[current][i]] {
				q = append(q, road2[current][i])
			}
		}
	}
}