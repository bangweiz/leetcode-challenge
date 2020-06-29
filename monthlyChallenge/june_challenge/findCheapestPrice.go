package june_challenge

import "fmt"

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if src == dst && K == 0 {
		return 0
	}
	data := make([][]int, K + 1)
	for i := range data {
		data[i] = make([]int, n)
	}
	edges := make([]map[int]int, n)
	for i := range edges {
		edges[i] = make(map[int]int)
	}
	for i := range flights {
		edges[flights[i][0]][flights[i][1]] = flights[i][2]
	}

	for key, value := range edges[src] {
		data[0][key] = value
	}

	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data[i]); j++ {
			data[i + 1][j] = data[i][j]
		}
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != 0 {
				for key, value := range edges[j] {
					if data[i + 1][key] == 0 {
						data[i + 1][key] = data[i][j] + value
					} else {
						data[i + 1][key] = min(data[i + 1][key], data[i][j] + value)
					}
				}
			}
		}
	}

	for _, v1 := range data {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	if data[K][dst] == 0 {
		return -1
	}
	return data[K][dst]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

