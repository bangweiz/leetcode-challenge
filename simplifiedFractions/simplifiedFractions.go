package simplifiedFractions

import "strconv"

func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	for i := 2; i <= n; i++ {
		res = append(res, "1/" + strconv.Itoa(i))
		for j := 2; j < i; j++ {
			if helper(i, j) {
				res = append(res, strconv.Itoa(j) + "/" + strconv.Itoa(i))
			}
		}
	}
	return res
}

func helper(n1, n2 int) bool {
	for n1 % n2 != 0 {
		n1, n2 = n2, n1 % n2
	}
	return n2 == 1
}