package maxPower

func maxPower(s string) int {
	b := []byte(s)
	if len(b) == 0 {
		return 0
	}
	max, temp, current := 1, 1, b[0]
	for i := 1; i < len(b); i++ {
		if b[i] == current {
			temp++
		} else {
			if temp > max {
				max = temp
			}
			temp = 1
			current = b[i]
		}
	}
	if temp > max {
		max = temp
	}
	return max
}
