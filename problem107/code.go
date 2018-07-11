package problem107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	arr := make([][]int, 0)

	return levelOrderBottomChildren(arr, 0, []*TreeNode{root})
}

func levelOrderBottomChildren(arr [][]int, level int, children []*TreeNode) [][]int {
	present := make([]int, 0)
	next := make([]*TreeNode, 0)
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
