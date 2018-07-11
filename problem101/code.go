package problem101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Right == nil {
		return false
	}

	if root.Left == nil && root.Right != nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	left := root.Left
	right := root.Right
	return left.Val == right.Val && isSymmetricChildren(left, right)
}

func isSymmetricChildren(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if right == nil {
		return false
	}

	return left.Val == right.Val && isSymmetricChildren(left.Left, right.Right) && isSymmetricChildren(left.Right, right.Left)
}
