package leetcode

// 读懂题意，是解题的关键
// https://leetcode-cn.com/problems/maximum-binary-tree-ii/
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root==nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		tmp := &TreeNode{Val: val, Left: root}
		root = tmp
	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
	}
	return root
}