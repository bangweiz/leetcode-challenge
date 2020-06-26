package june_challenge

func sumNumbers(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	deepFirst(root, &res, 0)
	return
}

func deepFirst(root *TreeNode, count *int, nums int) {
	nums = nums * 10 + root.Val
	if root.Right == nil && root.Left == nil {
		*count += nums
	} else {
		if root.Left != nil {
			deepFirst(root.Left, count, nums)
		}
		if root.Right != nil {
			deepFirst(root.Right, count, nums)
		}
	}
}
