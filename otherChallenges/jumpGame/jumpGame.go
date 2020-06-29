package jumpGame

func CanJump(nums []int) bool {
	zeroArr := make([]int, 0)
	for  i := 0; i < len(nums); i++ {
		if nums[i] == 0 && i != len(nums) - 1 {
			zeroArr = append(zeroArr, i)
		}
	}
	if len(zeroArr) == 0 {
		return true
	}

	j := len(nums)
	for i := len(zeroArr) - 1; i >= 0; i-- {
		if j < zeroArr[i] {
			continue
		}
		count := 1
		flag := false
		for j = zeroArr[i]; j >= 0; j-- {
			if nums[j] >= count {
				flag = true
				break
			}
			count++
		}
		if !flag {
			return false
		}
	}
	return true
}

func canJump(nums []int) bool {
	// TODO: complete one loop method
	k := 0
	for i := len(nums) - 2; i >= 0; i++ {
		if nums[i] == 0 {
			k = i
			break
		}
	}
	for i := k; i >= 0; i++ {
		count := 1
		if nums[i] >= count {
			count = 1
		} else {
			count++
		}
	}
	return true
}
