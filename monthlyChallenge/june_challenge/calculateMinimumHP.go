package june_challenge

func calculateMinimumHP(dungeon [][]int) int {
	res, l1, l2 := make([][]int, len(dungeon)), len(dungeon), len(dungeon[0])
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, l2)
	}
	
	res[l1 - 1][l2 - 1] = 1

	for i := l2 - 2; i >= 0; i-- {
		res[l1 - 1][i] = max(1, res[l1 - 1][i + 1] - dungeon[l1 - 1][i + 1])
	}

	for i := l1 - 2; i >= 0; i-- {
		res[i][l2 - 1] = max(1, res[i + 1][l2 - 1] - dungeon[i + 1][l2 - 1])
	}
	
	for i := l1 - 2; i >= 0; i-- {
		for j := l2 - 2; j >= 0; j-- {
			temp1, temp2 := max(1, res[i + 1][j] - dungeon[i + 1][j]), max(1, res[i][j + 1] - dungeon[i][j + 1])
			res[i][j] = min(temp1, temp2)
		}
	}
	return max(1, res[0][0] - dungeon[0][0])
 }

 func max(a, b int) int {
	 if a > b {
		 return a
	 }
	 return b
 }

 //func min(a,b int) int {
	//if a < b {
	//	return a
	//}
	//return b
 //}
