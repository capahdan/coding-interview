package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	leftTree := IsSameTree(p.Left, q.Left)
	rightTree := IsSameTree(p.Right, q.Right)

	return leftTree && rightTree
}
