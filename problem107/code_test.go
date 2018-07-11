package problem107

import (
	"testing"
)

func buildTree(vals []int) *TreeNode {
	if vals == nil || len(vals) == 0 {
		return nil
	}
	root := &TreeNode{}
	buildTreeNode(vals, 0, root)

	return root
}

func buildTreeNode(vals []int, position int, node *TreeNode) {
	if 2*position+2 >= len(vals) {
		return
	}
	if position == 0 {
		node.Val = vals[2*position]
	}
	left := vals[2*position+1]
	right := vals[2*position+2]
	if left != 0 {
		node.Left = &TreeNode{Val: left}
	}
	if right != 0 {
		node.Right = &TreeNode{Val: right}
	}
	buildTreeNode(vals, 2*position+1, node.Left)
	buildTreeNode(vals, 2*position+2, node.Right)
}

func Test1(t *testing.T) {
	exp := [][]int{{15, 7}, {9, 20}, {3}}
	res := levelOrderBottom(buildTree([]int{3, 9, 20, 0, 0, 15, 7}))

	flag := false
	for i, v := range exp {
		rV := res[i]
		iFlag := false
		for j, val := range v {
			if val != rV[j] {
				iFlag = true
				break
			}
		}
		if iFlag {
			flag = true
			break
		}
	}
	if flag {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}
