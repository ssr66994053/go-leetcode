package P257

import (
	"fmt"
	"github.com/ssr66994053/go-leetcode/model"
)

func binaryTreePaths(root *model.TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	arr := make([]string, 0)
	if root.Left != nil {
		for _, path := range binaryTreePaths(root.Left) {
			arr = append(arr, fmt.Sprintf("%d->%s", root.Val, path))
		}
	}
	if root.Right != nil {
		for _, path := range binaryTreePaths(root.Right) {
			arr = append(arr, fmt.Sprintf("%d->%s", root.Val, path))
		}
	}
	return arr
}
