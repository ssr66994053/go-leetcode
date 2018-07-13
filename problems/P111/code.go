package P111

import "github.com/ssr66994053/go-leetcode/model"

func minDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 && right != 0 {
		return right + 1
	}
	if left != 0 && right == 0 {
		return left + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}
