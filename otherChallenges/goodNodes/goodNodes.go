package goodNodes

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxStack, max, res := []int{-100000}, -100000, 0
	dfs(root, maxStack, &res, max)
	return res
}

func dfs(root *TreeNode, maxStack []int, res *int, max int) {
	if root.Val >= max {
		*res += 1
		maxStack, max = append(maxStack, root.Val), root.Val
	}

	if root.Left != nil {
		dfs(root.Left, maxStack, res, max)
	}

	if root.Right != nil {
		dfs(root.Right, maxStack, res, max)
	}

	if root.Val == max {
		maxStack = maxStack[:len(maxStack) - 1]
		max = maxStack[len(maxStack) - 1]
	}
}
