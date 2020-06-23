package june_challenge

func countNodes(root *TreeNode) int {
	count := 0
	if root != nil {
		s(root, &count)
	}
	return count
}

func s(root *TreeNode, c *int) {
	*c += 1
	if root.Left != nil {
		s(root.Left, c)
	}
	if root.Right != nil {
		s(root.Right, c)
	}
}
