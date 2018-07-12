package P104

import "github.com/ssr66994053/go-leetcode/model"

func maxDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
