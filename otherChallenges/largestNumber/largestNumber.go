package largestNumber

import "strconv"

func largestNumber(cost []int, target int) string {
	res := make([]string, target + 1)
	for j, v := range cost {
		if v < target + 1 {
			res[v] = strconv.Itoa(j + 1)
		}
	}

	for i := 1; i < target + 1; i++ {
		for j := 0; j < 9; j++ {
			if i - cost[j] > 0 && res[i - cost[j]] != "" && res[cost[j]] != "" {
				res[i] = compareStr(combineStr( res[i - cost[j]], res[cost[j]]), res[i])
			}
		}
	}
	if res[target] == "" {
		return "0"
	}
	return res[target]
}

func compareStr(s1, s2 string) string {
	b1, b2 := []byte(s1), []byte(s2)
	if len(b1) < len(b2) {
		return s2
	} else if len(b1) > len(b2) {
		return s1
	}

	i := 0
	for i < len(b1) && i < len(b2) {
		if b1[i] < b2[i] {
			return s2
		} else if b1[i] > b2[i] {
			return s1
		}
		i++
	}
	return s1
}

func combineStr(s1, s2 string) string {
	return compareStr(s1 + s2, s2 + s1)
}