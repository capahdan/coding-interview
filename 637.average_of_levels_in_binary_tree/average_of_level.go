package average_of_level

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelLength := len(queue)
		levelSum := 0

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		average := float64(levelSum) / float64(levelLength)
		result = append(result, average)
	}

	return result
}
