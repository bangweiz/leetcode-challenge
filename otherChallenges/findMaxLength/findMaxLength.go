package findMaxLength

func findMaxLength(nums []int) int {
	myMap := map[int]int{0: -1}
	count := 0
	max := 0
	for i, v := range nums {
		if v == 0 {
			count -= 1
		} else {
			count += 1
		}
		value, ok := myMap[count];
		if ok {
			if max < i - value {
				max = i - value
			}
		} else {
			myMap[count] = i
		}
	}
	return max
}
