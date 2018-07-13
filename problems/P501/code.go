package P501

import (
	"github.com/ssr66994053/go-leetcode/model"
)

func findMode(root *model.TreeNode) []int {
	m := make(map[int]int, 0)
	walkThrough(root, m)

	n := 0
	arr := make([]int, 0)
	for k, v := range m {
		if n == 0 {
			n = v
		}
		if n < v {
			n = v
			arr = make([]int, 0)
		}
		if n == v {
			arr = append(arr, k)
		}
	}

	return arr
}

func walkThrough(node *model.TreeNode, m map[int]int) {
	if node == nil {
		return
	}

	if v, ok := m[node.Val]; ok {
		m[node.Val] = v + 1
	} else {
		m[node.Val] = 1
	}

	walkThrough(node.Left, m)
	walkThrough(node.Right, m)
}
