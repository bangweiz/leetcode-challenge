package moveZeroes

func moveZeroes(nums []int)  {
	pointer := -1
	for i := range nums {
		if nums[i] == 0 {
			pointer = i
			break
		}
	}
	if pointer == -1 {
		return
	}
	for i := pointer + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[pointer] = nums[pointer], nums[i]
			pointer++
		}
	}
}
