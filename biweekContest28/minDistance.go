package biweekContest28

import "fmt"

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

type SubrectangleQueries struct {
	Grid [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}


func (s *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			s.Grid[i][j] = newValue
		}
	}
}


func (s *SubrectangleQueries) GetValue(row int, col int) int {
	return s.Grid[row][col]
}

func minSumOfLengths(arr []int, target int) int {
	prefix, sufix := make([]int, len(arr)), make([]int, len(arr))
	for i := range arr {
		temp, count, flag := target, 0, false
		for j := i; j >= 0; j-- {
			temp -= arr[j]
			count++
			if temp == 0 {
				prefix[i], flag = count, true
				break
			} else if temp < 0 {
				break
			}
		}
		if !flag {
			prefix[i] = 0
		}
		temp, count, flag = target, 0, false
		for j := i; j < len(arr); j++ {
			temp -= arr[j]
			count++
			if temp == 0 {
				sufix[i], flag = count, true
				break
			} else if temp < 0 {
				break
			}
		}
		if !flag {
			sufix[i] = 0
		}
	}

	res := 999999999
	for i := 0; i < len(prefix) - 1; i++ {
		if prefix[i] != 0 && sufix[i + 1] != 0 {
			res = min(res, prefix[i] + sufix[i + 1])
		}
	}
	if res == 999999999 {
		return -1
	}
	return res
}


func minDistance(houses []int, k int) int {
	if len(houses) == k {
		return 0
	}
	res := make([][]int, k)
	for i := range res {
		res[i] = make([]int, len(houses))
	}
	houses = mergeSort(houses)

	res[0][1] = houses[1] - houses[0]
	for i := 2; i < len(res[0]); i++ {
		var n int
		if i % 2 == 0 {
			n = i / 2
		} else {
			n = i / 2 + 1
		}
		for j := 0; j <= i; j++ {
			if j < n {
				res[0][i] += houses[n] - houses[j]
			} else {
				res[0][i] += houses[j] - houses[n]
			}
		}
	}

	for i := 1; i < k; i++ {
		for j := i + 1; j < len(res[i]); j++ {
			res[i][j] = res[i - 1][j - 1]
			for g := j - 1; g >= i; g-- {
				sum, n := 0, 0

				if (j - g) % 2 == 0 {
					n = (j - g) / 2 + g
				} else {
					n = (j - g) / 2 + 1 + g
				}

				for l := g; l <= j; l++ {
					if l < n {
						sum += houses[n] - houses[l]
					} else {
						sum += houses[l] - houses[n]
					}
				}
				res[i][j] = min(res[i][j], res[i - 1][g - 1] + sum)
			}
		}
	}

	for _, v1 := range res {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	return res[k - 1][len(houses) - 1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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