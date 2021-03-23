package leetcode

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	if isBalanced(root.Left)==false || isBalanced(root.Right)==false{
		return false
	}
	left,right := Height(root.Left), Height(root.Right)
	if left-right>1 || left-right< -1 {
		return false
	}
	return true
}



func  Height(n *TreeNode) int {
	if n==nil {
		return 0
	}
	height := 1
	left,right := Height(n.Left), Height(n.Right)

	if left>right {
		height += left
	} else {
		height += right
	}
	return height
}

