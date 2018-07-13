package P107

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func levelOrderBottom(root *model.TreeNode) [][]int {
	arr := make([][]int, 0)

	if root == nil {
		return [][]int{}
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	left := levelOrderBottom(root.Left)
	right := levelOrderBottom(root.Right)
	length := len(left)
	if len(right) > len(left) {
		length = len(right)
	}
	for i := 0; i < length; i++ {
		a := make([]int, 0)
		if i-length+len(left) >= 0 {
			a = append(a, left[i-length+len(left)]...)
		}
		if i-length+len(right) >= 0 {
			a = append(a, right[i-length+len(right)]...)
		}
		arr = append(arr, a)
	}

	arr = append(arr, []int{root.Val})

	return arr
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
