package sorted_arr_to_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return SortedArrayHelper(nums, 0, len(nums)-1)
}

func SortedArrayHelper(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayHelper(nums, start, mid-1)
	root.Right = SortedArrayHelper(nums, mid+1, end)
	return root
}
