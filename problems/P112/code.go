package P112

import "github.com/ssr66994053/go-leetcode/model"

func hasPathSum(root *model.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	left := sum - root.Val
	if root.Left == nil && root.Right == nil && left == 0 {
		return true
	}

	if root.Left != nil {
		leftRes := hasPathSum(root.Left, left)
		if leftRes {
			return true
		}
	}
	if root.Right != nil {
		rightRes := hasPathSum(root.Right, left)
		if rightRes {
			return true
		}
	}

	return false
}
