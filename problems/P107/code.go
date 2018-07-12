package P107

import "github.com/ssr66994053/go-leetcode/model"

func levelOrderBottom(root *model.TreeNode) [][]int {
	arr := make([][]int, 0)

	return levelOrderBottomChildren(arr, 0, []*model.TreeNode{root})
}

func levelOrderBottomChildren(arr [][]int, level int, children []*model.TreeNode) [][]int {
	present := make([]int, 0)
	next := make([]*model.TreeNode, 0)
	for _, c := range children {
		if c != nil {
			present = append(present, c.Val)
			if c.Left != nil {
				next = append(next, c.Left)
			}
			if c.Right != nil {
				next = append(next, c.Right)
			}
		}
	}
	if len(present) == 0 {
		return arr
	}

	arr = levelOrderBottomChildren(arr, level+1, next)

	return append(arr, present)
}
