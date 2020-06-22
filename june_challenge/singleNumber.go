package june_challenge

func singleNumber(nums []int) int {
	hmap := make(map[int]int, 0)
	for _, v := range nums {
		value, ok := hmap[v]
		if ok {
			if value == 1 {
				hmap[v] = 2
			} else {
				delete(hmap, v)
			}
		} else {
			hmap[v] = 1
		}
	}
	for key, _ := range hmap {
		return key
	}
	return 0
}
