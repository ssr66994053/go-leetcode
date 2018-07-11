package problem101

import "testing"

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
	exp := true
	res := isSymmetric(buildTree([]int{1, 2, 2, 3, 4, 4, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := false
	res := isSymmetric(buildTree([]int{1, 2, 2, 0, 3, 0, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := true
	res := isSymmetric(buildTree([]int{1, 2, 2, 0, 3, 3, 0}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
