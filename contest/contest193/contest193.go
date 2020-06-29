package contest193

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i - 1]
	}
	return nums
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	myMap := make(map[int]int, 0)
	for _, v := range arr {
		v1, ok := myMap[v]
		if ok {
			myMap[v] = v1 + 1
		} else {
			myMap[v] = 1
		}
	}

	res := make([]int, 0)

	for _, value := range myMap {
		res = append(res, value)
	}
	res = mergeSort(res)
	count := len(res)
	for _, v := range res {
		k -= v
		if k > 0 {
			count--
		} else if k == 0 {
			return count - 1
		} else {
			return count
		}
	}
	return 0
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

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m * k {
		return -1
	}
	min, max := 1, bloomDay[0]
	for _, v := range bloomDay {
		if v > max {
			max = v
		}
	}
	bloomArr, mid := make([]int, 0), 0
	for min < max {
		bloomArr = []int{}
		mid = (min + max) / 2
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= mid {
				bloomArr = append(bloomArr, i)
			}
		}
		res := check(bloomArr, k)
		if res < m {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func check(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	count, pointer := 0, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == 1 + arr[i- 1] {
			pointer++
		} else {
			count += pointer / k
			pointer = 1
		}
	}
	return count + pointer / k
}

type TreeAncestor struct {
	parent []int
}


func Constructor(n int, parent []int) TreeAncestor {
	return TreeAncestor{parent}
}


func (t *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := 1; i <= k; i++ {
		if node < 0 {
			return -1
		}
		node = t.parent[node]
	}
	return node
}
