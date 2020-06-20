package june_challenge

func longestDupSubstring(S string) string {
	b, i := []byte(S), 0
	hash, power := make([]int, len(b) + 1), make([]int, len(b) + 1)
	power[0] = 1
	for i := 0; i < len(b); i++ {
		hash[i + 1] = (29 * hash[i] + int(b[i] - 'a')) % 1000000007
		power[i + 1] = power[i] * 29 % 1000000007
	}
	left, right, mid := 0, len(b) - 1, 0
	for left < right {
		mid = (left + right) / 2 + 1
		ok, index := rabinKarp(b, mid, 1000000007, hash, power)
		if ok {
			i = index
			left = mid
		} else {
			right = mid - 1
		}
	}
	return string(b[i: i + left])
}

func rabinKarp(b []byte, n, mod int, hash, power []int) (bool, int) {
	hmap := make(map[int][]int, 0)
	for i := n - 1; i < len(b); i++ {
		temp := (hash[i + 1] - power[n] * hash[i + 1 - n] % mod + mod) % mod
		if v, ok := hmap[temp]; ok {
			for _, index := range v {
				if collisions(b, i - n + 1, index, n) {
					hmap[temp] = append(hmap[temp], i - n + 1)
				} else {
					return true, i - n + 1
				}
			}
		} else {
			hmap[temp] = []int{i - n + 1}
		}
	}
	return false, -1
}

func collisions(b []byte, j, k, n int) (flag bool) {
	for i := 0; i < n; i++ {
		if b[j + i] != b[k + i] {
			return true
		}
	}
	return
}