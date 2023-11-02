package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {

	return root == nil || IsSymmetricHelp(root.Left, root.Right)
}

func IsSymmetricHelp(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}

	left := IsSymmetricHelp(p.Left, q.Right)
	right := IsSymmetricHelp(p.Right, q.Left)

	return left && right
}
