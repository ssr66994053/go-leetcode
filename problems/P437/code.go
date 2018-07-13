package P437

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func pathSum(root *model.TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return findPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func findPath(node *model.TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	m := 0
	n := sum - node.Val
	if n == 0 {
		m = m + 1
	}
	m = m + findPath(node.Left, n) + findPath(node.Right, n)

	return m
}
