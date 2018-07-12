package P226

import "github.com/ssr66994053/go-leetcode/model"

func invertTree(root *model.TreeNode) *model.TreeNode {
	if root != nil {
		tmp := invertTree(root.Right)
		root.Right = invertTree(root.Left)
		root.Left = tmp
	}

	return root
}
