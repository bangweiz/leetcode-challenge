package june_challenge

import "math/rand"

type Solution struct {
	p []int
}

func constructor(w []int) Solution {
	sum, p := 0, make([]int, len(w))
	for i, v := range w {
		sum += v
		p[i] = sum
	}

	return Solution{p}
}


func (s *Solution) PickIndex() int {
	n := rand.Intn(s.p[len(s.p) - 1])
	left, right := 0, len(s.p) - 1
	for left <= right {
		mid := (left + right) / 2
		if s.p[mid] == n {
			return mid + 1
		} else if s.p[mid] > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
