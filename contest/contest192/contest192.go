package contest192

import (
	"math"
)

func shuffle(nums []int, n int) []int {
	if len(nums) <= 2 {
		return nums
	}
	arr1, arr2 := nums[:n], nums[n:]
	res := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		res = append(res, arr1[i])
		res = append(res, arr2[i])
	}
	return res
}

func getStrongest(arr []int, k int) []int {
	arr = mergeSort(arr)
	m := arr[(len(arr) - 1) / 2]
	p1, p2, res := 0, len(arr) - 1, make([]int, 0)
	for i := 0; i < k; i++ {
		if math.Abs(float64(arr[p1]) - float64(m)) > math.Abs(float64(arr[p2]) - float64(m)) {
			res = append(res, arr[p1])
			p1++
		} else {
			res = append(res, arr[p2])
			p2--
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

type BrowserHistory struct {
	urls []string
	current int
}


func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}


func (b *BrowserHistory) Visit(url string)  {
	data := b.urls[:b.current + 1]
	b.urls = data
	b.urls = append(b.urls, url)
	b.current = len(b.urls) - 1
}


func (b *BrowserHistory) Back(steps int) string {
	target := b.current - steps
	if target <= 0 {
		b.current = 0
		return b.urls[0]
	}
	b.current = target
	return b.urls[target]
}


func (b *BrowserHistory) Forward(steps int) string {
	target := b.current + steps
	if target >= len(b.urls) {
		b.current = len(b.urls) - 1
		return b.urls[len(b.urls) - 1]
	}
	b.current = target
	return b.urls[target]
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	res := make([][][]int, m + 1)
	for i := range res {
		res[i] = make([][]int, n + 1)
		for j := range res[i] {
			res[i][j] = make([]int, target + 1)
		}
	}
	for i := 1; i < len(res); i++ {
		for k := 1; k < target + 1; k++ {
			for j := 1; j < len(res[i]); j++ {
				if k > i {
					res[i][j][k] = 999999999
				} else if houses[i - 1] != 0 {
					if houses[i - 1] != j {
						res[i][j][k] = 999999999
					} else {
						temp := res[i - 1][j][k]
						for index := 0; index < len(cost[0]); index++ {
							if index != j - 1 && k - 1 != 0 {
								temp = min(temp, res[i - 1][index + 1][k - 1])
							}
						}
						res[i][j][k] = temp
					}
				} else {
					temp := res[i - 1][j][k] + cost[i - 1][j - 1]
					for index := 0; index < len(cost[0]); index++ {
						if index + 1 != j && k - 1 != 0 {
							temp = min(temp, cost[i - 1][j - 1] + res[i - 1][index + 1][k - 1])
						}
					}
					res[i][j][k] = temp
				}
			}
		}
	}
	max := 999999999
	for i := 1; i <= n; i++ {
		max = min(max, res[m][i][target])
	}
	if max == 999999999 {
		 return -1
	}
	return max
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}