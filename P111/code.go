package P111

import "github.com/ssr66994053/go-leetcode/model"

func minDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepthChild(root.Left, 1)
	right := minDepthChild(root.Right, 1)
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

func minDepthChild(node *model.TreeNode, depth int) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return depth
	}

	left := minDepthChild(node.Left, depth+1)
	right := minDepthChild(node.Right, depth+1)
	if left == 0 && right != 0 {
		return right
	}
	if left != 0 && right == 0 {
		return left
	}
	if left < right {
		return left
	}
	return right
}
