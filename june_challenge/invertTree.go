package june_challenge

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	bfs(q)
	return root
}

func bfs(q []*TreeNode) {
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		current.Left, current.Right = current.Right, current.Left
		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}
	}
}