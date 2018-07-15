package P543

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func diameterOfBinaryTree(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	arr := walkThrough(root, []int{})
	for _, d := range arr {
		if d > max {
			max = d
		}
	}

	return max
}

func walkThrough(node *model.TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}

	l, r := 0, 0
	if node.Left != nil {
		l = maxDepth(node.Left)
		arr = walkThrough(node.Left, arr)
	}
	if node.Right != nil {
		r = maxDepth(node.Right)
		arr = walkThrough(node.Right, arr)
	}

	return append(arr, l+r)
}

func maxDepth(node *model.TreeNode) int {
	if node == nil {
		return 0
	}

	l := 1 + maxDepth(node.Left)
	r := 1 + maxDepth(node.Right)

	if l > r {
		return l
	}
	return r
}
