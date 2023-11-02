package invert_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	left := InvertTree(root.Left)
	right := InvertTree(root.Right)

	return &TreeNode{
		Val:   root.Val,
		Left:  right,
		Right: left,
	}
}
