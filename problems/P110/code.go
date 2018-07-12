package P110

import "github.com/ssr66994053/go-leetcode/model"

func isBalanced(root *model.TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) {
		return false
	}
	if !isBalanced(root.Right) {
		return false
	}

	leftDepth := childrenDepth(root.Left, 1)
	rightDepth := childrenDepth(root.Right, 1)
	if leftDepth-rightDepth >= -1 && leftDepth-rightDepth <= 1 {
		return true
	}

	return false
}

func childrenDepth(node *model.TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	leftDepth := childrenDepth(node.Left, depth+1)
	rightDepth := childrenDepth(node.Right, depth+1)
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}
