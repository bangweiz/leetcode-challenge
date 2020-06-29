package groupAnagrams

func GroupAnagrams(str []string) [][]string {
	myMap := make(map[int][]string, len(str))
	for _, v := range str {
		index := hashFunc(v)
		if i, ok := myMap[index]; ok {
			myMap[index] = append(i, v)
		} else {
			myMap[index] = make([]string, 1)
			myMap[index][0] = v
		}
	}
	res := make([][]string, 0)
	for _, v := range myMap {
		res = append(res, v)
	}
	return res
}

func hashFunc(str string) (res int) {
	bytes := []byte(str)
	res = 1
	for _, v := range bytes {
		res *= value(v)
	}
	return
}

func value(char byte) (res int) {
	switch char {
	case 122: return 101
	case 121: return 2
	case 120: return 3
	case 119: return 5
	case 118: return 7
	case 117: return 11
	case 116: return 13
	case 115: return 17
	case 114: return 19
	case 113: return 23
	case 112: return 29
	case 111: return 31
	case 110: return 37
	case 109: return 41
	case 108: return 43
	case 107: return 47
	case 106: return 53
	case 105: return 59
	case 104: return 61
	case 103: return 67
	case 102: return 71
	case 101: return 73
	case 100: return 79
	case 99: return 83
	case 98: return 89
	case 97: return 97
	default: return 1
	}
}
