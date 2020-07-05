package contest196

func canMakeArithmeticProgression(arr []int) bool {
	arr = mergeSort(arr)
	res := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i] - arr[i - 1] != res {
			return false
		}
	}
	return true
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

func getLastMoment(n int, left []int, right []int) int {
	res := 0
	for _, v := range left {
		if res < v {
			res = v
		}
	}
	for _, v := range right {
		if res < (n - v) {
			res = n - v
		}
	}
	return res
}

func numSubmat(mat [][]int) (res int) {
	dp := make([][]int, len(mat))
	for i := range mat {
		dp[i] = make([]int, len(mat[i]));
	}

	dp[0][0], res = mat[0][0], mat[0][0]

	for i := 1; i < len(mat[0]); i++ {
		if mat[0][i] == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1 + dp[0][i - 1]
		}
		res += dp[0][i]
	}

	for i := 1; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				height, width := 0, 0
				for i >= height + 1 && dp[i - 1 - height][j] != 0 {
					height++
				}
				for j >= width + 1 && dp[i][j - 1 - width] != 0 {
					width++
				}
				dp[i][j] += 1 + height + width
				if height > 0 && width > 0 {
					dp[i][j] += count(mat, i, j, height, width)
				}
				res += dp[i][j]
			}
		}
	}
	return
}

func count(mat [][]int, i, j, height, width int) (res int) {
	for h := 1; h <= height; h++ {
		for w := 1; w <= width; w++ {
			if mat[i -h][j - w] == 0 {
				width = w - 1
				break
			}
		}
		res += width
	}
	return
}

func minInteger(num string, k int) string {
	b, start := []byte(num), 0
	for k > 0 {
		index, count, c := start, 0, 0
		for i := start + 1; i < len(b) && c < k; i++ {
			c++
			if b[index] > b[i] {
				index, count = i, c
			}
		}
		if count == 0 {
			start++
			continue
		}
		if start >= len(b) {
			break;
		}
		t := make([]byte, 0)
		t = append(t, b[:index - count]...)
		t = append(t, b[index])
		t = append(t, b[index - count : index]...)
		t = append(t, b[index + 1:]...)
		b, k, start = t, k - count, start + 1
	}
	return string(b)
}