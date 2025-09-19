package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	need := targetSum
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		return need == root.Val
	} else {
		left := hasPathSum(root.Left, need-root.Val)
		right := hasPathSum(root.Right, need-root.Val)
		return left || right
	}
}
