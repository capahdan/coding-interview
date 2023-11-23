package count_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDept := CountNodes(root.Left)
	rightDept := CountNodes(root.Right)

	return leftDept + rightDept + 1
}
