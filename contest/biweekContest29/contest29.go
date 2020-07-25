package biweekContest29

func average(salary []int) float64 {
	min, max, res := salary[0], salary[0], salary[0]
	for i := 1; i < len(salary); i++ {
		res += salary[i]
		if salary[i] < min {
			min = salary[i]
		} else if salary[i] > max {
			max = salary[i]
		}
	}
	res = res - min - max
	return float64(res) / float64(len(salary) - 2)
}

func kthFactor(n int, k int) int {
	res := []int{1}
	for i := 2; i <= n; i++ {
		if n % i == 0 {
			res = append(res, i)
		}
	}
	if len(res) < k {
		return -1
	} else {
		return res[k - 1]
	}
}

func longestSubarray(nums []int) int {
	prev, current, res := 0, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			if i + 1 != len(nums) {
				if current + prev > res {
					res = current + prev
				}
				if nums[i + 1] == 1 {
					prev = current
				}
				current = 0
			}
		} else {
			current++
		}
	}

	if current + prev == len(nums) {
		return len(nums) - 1
	}

	if current + prev > res {
		return current + prev
	}

	return res
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {

}