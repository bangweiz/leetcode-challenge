package june_challenge

func ReconstructQueue(people [][]int) [][]int {
	people = mergeSort(people)
	res := make([][]int, 0)
	for _, v:= range people {
		if v[1] == len(res) {
			res = append(res, v)
		} else {
			res = append(res[:v[1]], append([][]int{v}, res[v[1]:]...)...)

		}
	}
	return res
}

func mergeSort(nums [][]int) [][]int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(a, b [][]int) [][]int {
	c, i, j, index := make([][]int, len(a) + len(b)), 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i][0] > b[j][0] {
			c[index] = a[i]
			i++
		} else if a[i][0] < b[j][0] {
			c[index] = b[j]
			j++
		} else {
			if a[i][1] < b[j][1] {
				c[index] = a[i]
				i++
			} else {
				c[index] = b[j]
				j++
			}
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