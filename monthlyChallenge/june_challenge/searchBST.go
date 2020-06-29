package june_challenge

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if res := searchBST(root.Left, val); res != nil {
		return res
	}
	if res := searchBST(root.Right, val); res != nil {
		return res
	}
	return nil
}
