package new21Game

func new21Game(N int, K int, W int) float64 {
	if W + K <= N || K == 0 {
		return 1
	}
	k := float64(1) / float64(W)
	p := make([]float64, N + 1)
	sum, res := float64(0), float64(0)

	for i := 1; i < N + 1; i++ {
		var m float64
		if i <= W {
			m = k
		} else {
			m = 0
		}
		p[i] = sum * k + m
		if i < K {
			sum += p[i]
		}

		if i > W {
			sum -= p[i - W]
		}

		if i >= K {
			res += p[i]
		}
	}

	return res
}

