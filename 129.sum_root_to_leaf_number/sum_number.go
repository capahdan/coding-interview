package sum_number

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumber(root *TreeNode) int {

	return sumNumberHelper(root, 0)
}

// yang pasti kita harus mengunjungi setiap node
func sumNumberHelper(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return sum*10 + root.Val
	}

	return sumNumberHelper(root.Left, sum*10+root.Val) + sumNumberHelper(root.Right, sum*10+root.Val)
}
