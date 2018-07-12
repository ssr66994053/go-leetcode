package model

import (
	"encoding/json"
)

type TreeNode struct {
	Val   int       `json:val`
	Left  *TreeNode `json:left`
	Right *TreeNode `json:right`
}

func (t *TreeNode) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (t *TreeNode) StringArray() []int {
	arr := make([]int, 0)
	pipe := make(chan int, 0)
	done := make(chan bool, 0)
	go func() {
		flat(t, pipe)
		done <- true
	}()

	for flag := true; flag; {
		select {
		case v := <-pipe:
			arr = append(arr, v)
		case <-done:
			flag = false
		}
	}
	return arr
}

func flat(t *TreeNode, pipe chan int) {
	if t.Left != nil {
		flat(t.Left, pipe)
	}
	pipe <- t.Val
	if t.Right != nil {
		flat(t.Right, pipe)
	}
}

func BuildTree(vals []int) *TreeNode {
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
	if left != 999 {
		node.Left = &TreeNode{Val: left}
		buildTreeNode(vals, 2*position+1, node.Left)
	}
	if right != 999 {
		node.Right = &TreeNode{Val: right}
		buildTreeNode(vals, 2*position+2, node.Right)
	}
}

func CheckTree(exp *TreeNode, tar *TreeNode) bool {
	if exp == nil && tar == nil {
		return true
	}
	if exp == nil {
		return false
	}
	if tar == nil {
		return false
	}
	if exp.Val != tar.Val {
		return false
	}
	return CheckTree(exp.Left, tar.Left) && CheckTree(exp.Right, exp.Right)
}
