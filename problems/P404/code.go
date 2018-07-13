package P404

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func sumOfLeftLeaves(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	n := 0
	if root.Left != nil {
		root.Val = 0
		left := root.Left
		if left.Left == nil && left.Right == nil {
			n = left.Val
		} else {
			n = sumOfLeftLeaves(left)
		}
	}
	if root.Right != nil {
		root.Val = 0
		right := root.Right
		if right.Left != nil || right.Right != nil {
			n = n + sumOfLeftLeaves(right)
		}
	}

	return n
}
