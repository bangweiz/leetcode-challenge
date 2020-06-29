package peopleIndexes

func peopleIndexes(favoriteCompanies [][]string) []int {
	res := make([]int, 0)
	myMap := make([]map[string]int, len(favoriteCompanies))
	for i := range favoriteCompanies {
		myMap[i] = make(map[string]int, 0)
		for j := range favoriteCompanies[i] {
			myMap[i][favoriteCompanies[i][j]] = 1
		}
	}

	for i := range favoriteCompanies {
		flag2 := true
		for j := range favoriteCompanies {
			if i != j && len(favoriteCompanies[i]) <= len(favoriteCompanies[j]) {
				flag := true
				for _, v := range favoriteCompanies[i] {
					if _, ok := myMap[j][v]; !ok {
						flag = false
						break
					}
				}
				if flag {
					flag2 = false
					break
				}
			}
		}
		if flag2 {
			res = append(res, i)
		}
	}

	return res
}
