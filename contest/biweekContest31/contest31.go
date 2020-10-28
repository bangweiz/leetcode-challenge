package biweekContest31

func countOdds(low int, high int) int {
	if low % 2 == 0 && high % 2 == 0 {
		return (high - low) / 2
	} else {
		return (high - low) / 2 + 1
	}
}

func numOfSubarrays(arr []int) int {
	res, subArrays := make([]int, len(arr)), [2]int{0, 0}
	if arr[0] % 2 == 0 {
		res[0], subArrays[0] = 0, 1
	} else {
		res[0], subArrays[1] = 1, 1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			res[i] = res[i - 1] + subArrays[1]
			subArrays[0] += 1
		} else {
			res[i] = res[i - 1] + subArrays[0] + 1
			subArrays[0], subArrays[1] = subArrays[1], subArrays[0] + 1
		}
	}
	return res[len(res) - 1] % 1000000007
}

func numSplits(s string) int {
	leftMap, rightMap, b, res := make(map[byte]int), make(map[byte]int), []byte(s), 0
	for _, v := range b {
		if val, ok := rightMap[v]; ok {
			rightMap[v] = val + 1
		} else {
			rightMap[v] = 1
		}
	}

	for _, v := range b {
		if num, _ := rightMap[v]; num == 1 {
			delete(rightMap, v)
		} else {
			rightMap[v] = num - 1
		}
		leftMap[v] = 1

		if len(leftMap) == len(rightMap) {
			res++
		}
	}

	return res
}

func minNumberOperations(target []int) int {
	res := make([]int, len(target))
	res[0] = target[0]
	for i := 1; i < len(res); i++ {
		if target[i] <= target[i - 1] {
			res[i] = res[i - 1]
		} else {
			res[i] = res[i - 1] + target[i] - target[i - 1]
		}
	}
	return res[len(res) - 1]
}