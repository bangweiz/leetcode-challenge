package contest194

import (
	"strconv"
)

func xorOperation(n int, start int) int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = start + 2 * i
	}
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = res ^ arr[i]
	}
	return res
}

func getFolderNames(names []string) []string {
	hashMap := make(map[string]int, 0)
	for i, v := range names {
		value, ok := hashMap[v]
		if ok {
			for true {
				temp := v + "(" + strconv.Itoa(value + 1) + ")"
				_, ok := hashMap[temp]
				if ok {
					value++
				} else {
					names[i] = temp
					hashMap[temp] = 0
					hashMap[v] = value + 1
					break
				}
			}
		} else {
			hashMap[v] = 0
		}
	}
	return names
}

func avoidFlood(rains []int) []int {
	start := 0
	for i := range rains {
		if rains[i] != 0 {
			start = i
			break
		}
	}
	hmap, zeroIndex := make(map[int]int, 0), make([]int, 0)

	for i := start; i < len(rains); i++ {
		if rains[i] == 0 {
			zeroIndex = append(zeroIndex, i)
		} else {
			if v, ok := hmap[rains[i]]; ok {
				if len(zeroIndex) == 0 {
					return make([]int, 0)
				} else {
					index := -1
					for j := 0; j < len(zeroIndex); j++ {
						if zeroIndex[j] > v {
							index = zeroIndex[j]
							zeroIndex = append(zeroIndex[: j], zeroIndex[j + 1:]...)
							break
						}
					}
					if index != -1 {
						rains[index] = rains[i]
						hmap[rains[i]] = i
						rains[i] = -1
					} else {
						return make([]int, 0)
					}
				}
			} else {
				hmap[rains[i]] = i
				rains[i] = -1
			}
		}
	}

	for i := range rains {
		if rains[i] == 0 {
			rains[i] = 1
		}
	}

	return rains
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	
}
