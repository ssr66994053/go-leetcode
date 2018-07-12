package P108

import "github.com/ssr66994053/go-leetcode/model"

func sortedArrayToBST(nums []int) *model.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	rootIdx := int(length / 2)
	node := &model.TreeNode{Val: nums[rootIdx]}
	if rootIdx > 0 {
		left := nums[:rootIdx]
		node.Left = sortedArrayToBST(left)
	}
	if rootIdx+1 < length {
		right := nums[rootIdx+1:]
		node.Right = sortedArrayToBST(right)
	}

	return node
}
