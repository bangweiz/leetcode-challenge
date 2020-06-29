package isHappy

func isHappy(n int) bool {
	myMap := make(map[int]int, 0)
	myMap[n] = 1
	for true {
		n = cal(n)
		if n == 1 {
			return true
		}
		if _, ok := myMap[n]; !ok {
			myMap[n] = 0
		} else {
			return false
		}
	}
	return false
}

func cal(n int) int {
	res := 0
	for n != 0 {
		num := n % 10
		n = n / 10
		if num != 0 {
			res += num * num
		}
	}
	return res
}