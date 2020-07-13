package biweekContest30

import (
	"math"
	"strings"
)

func reformatDate(date string) (res string) {
	d := strings.Split(date, " ")
	res = d[2]
	switch d[1] {
	case "Jan":
		res += "-01"
		break
	case "Feb":
		res += "-02"
		break
	case "Mar":
		res += "-03"
		break
	case "Apr":
		res += "-04"
		break
	case "May":
		res += "-05"
		break
	case "Jun":
		res += "-06"
		break
	case "Jul":
		res += "-07"
		break
	case "Aug":
		res += "-08"
		break
	case "Sep":
		res += "-09"
		break
	case "Oct":
		res += "-10"
		break
	case "Nov":
		res += "-11"
		break
	default:
		res += "-12"
		break
	}

	b := []byte(d[0])
	if len(b) == 3 {
		res += "-0" + string(b[0])
	} else {
		res += "-" + string(b[:2])
	}

	return
}

func rangeSum(nums []int, n int, left int, right int) (res int) {
	mod, data := 1000000007, make([]int, 0)
	for i := 0; i < n; i++ {
		temp := 0
		for j := i; j < n; j++ {
			temp = (nums[j] + temp) % mod
			data = append(data, temp)
		}
	}

	data = mergeSort(data)

	for i := left - 1; i < right; i++ {
		res = (res + data[i]) % mod
	}
	return
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

func minDifference(nums []int) (res int) {
	if len(nums) <= 4 {
		res = 0
	} else {
		res = math.MaxInt32
		nums = mergeSort(nums)
		for i := 0; i <= 3; i++ {
			temp := nums[len(nums) - 1 - i] - nums[3 - i]
			if temp < res {
				res = temp
			}
		}
	}
	return
}

func winnerSquareGame(n int) bool {
	dp, pows := make([]bool, n + 1), []int{1}
	dp[0] = false
	dp[1] = true
	for pows[len(pows) - 1] < n {
		pows = append(pows, (len(pows) + 1) * (len(pows) + 1))
	}

	for i := 2; i < n; i++ {
		dp[i] = check(dp, pows, i)
	}
	return dp[n]
}

func check(dp []bool, pows []int, i int) bool {
	for j := 0; pows[j] <= i; j++ {
		temp := i - pows[j]
		if temp == 0 {
			return true
		}
		if !dp[temp] {
			return true
		}
	}
	return false
}