package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := []int{3, 9, 20, null, null, 15, 7}
	q := []int{1, 2, 1}
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

func mirror(L, R) bool {
	if L == nil && R == nil {
		return true
	} else if L == nil || R == nil {
		return false
	} else if L.Val != R.Val {
		return false
	}
	return mirror(L.Left, R.Right) && mirror(L.Right, R.Left)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	result := 1 + max(leftDepth, rightDepth)
	return result
}
