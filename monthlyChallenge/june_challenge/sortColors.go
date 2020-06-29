package june_challenge

func sortColors(nums []int)  {
	index := make([]int, 3)
	for _, v := range nums {
		index[v]++
	}
	count := 0
	for j, v := range index {
		for i := 1; i <= v; i++ {
			nums[count] = j
			count++
		}
	}
}
