package max_depth_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDept := MaxDepth(root.Left)
	rightDept := MaxDepth(root.Right)

	if leftDept > rightDept {
		return leftDept + 1
	} else {
		return rightDept + 1
	}
}
