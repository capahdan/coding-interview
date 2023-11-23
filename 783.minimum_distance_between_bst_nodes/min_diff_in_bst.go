package min_dif_in_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetMinimumDifference(root *TreeNode) int {
	_, _, diff := helper(root)
	return diff
}

func helper(node *TreeNode) (int, int, int) {
	min, max, diff := node.Val, node.Val, math.MaxInt32
	if node.Left != nil {
		lMin, lMax, lDiff := helper(node.Left)
		min = lMin
		if tempDiff := node.Val - lMax; tempDiff < diff {
			diff = tempDiff
		}
		if lDiff < diff {
			diff = lDiff
		}
	}
	if node.Right != nil {
		rMin, rMax, rDiff := helper(node.Right)
		max = rMax
		if tempDiff := rMin - node.Val; tempDiff < diff {
			diff = tempDiff
		}
		if rDiff < diff {
			diff = rDiff
		}
	}
	return min, max, diff
}
