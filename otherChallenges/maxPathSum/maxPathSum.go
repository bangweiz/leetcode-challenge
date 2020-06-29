package maxPathSum

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var max int
var min int

func MaxPathSum(root *TreeNode) int {
	max = -4294967295
	min = -4294967295
	_, res2 := dfs(root)
	if min < 0 {
		return min
	}
	if res2 > max {
		return res2
	}
	return max
}

func dfs(node *TreeNode) (res1, res2 int) {
	var l1, l2, r1, r2 int
	if node.Left != nil {
		l1, l2 = dfs(node.Left)
		if l2 > max {
			max = l2
		}
	}
	if node.Right != nil {
		r1, r2 = dfs(node.Right)
		if r2 > max {
			max = r2
		}
	}
	if node.Val > min {
		min = node.Val
	}
	res2 = l1 + r1 + node.Val
	if l1 > r1 {
		res1 = node.Val + l1
	} else {
		res1 = node.Val + r1
	}
	if res1 < 0 {
		return 0, 0
	}
	return
}