package P538

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func convertBST(root *model.TreeNode) *model.TreeNode {
	walkThrough(root, 0)

	return root
}

func walkThrough(node *model.TreeNode, pre int) int {
	if node == nil {
		return 0
	}

	if node.Right != nil {
		pre = walkThrough(node.Right, pre)
	}

	pre = node.Val + pre
	node.Val = pre

	if node.Left != nil {
		pre = walkThrough(node.Left, pre)
	}

	return pre
}
