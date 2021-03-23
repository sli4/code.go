package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root==nil {
		return []int{}
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result =append(result, inorderTraversal(root.Right)...)
	return result
}


// https://leetcode-cn.com/problems/increasing-order-search-tree/
func increasingBST(root *TreeNode) *TreeNode {
	sl := inorderTraversal(root)
	var nroot,npre *TreeNode
	for _, v := range sl {
		tmp := &TreeNode{Val: v}
		if nroot==nil {
			nroot = tmp
			npre = tmp
		} else {
			npre.Right = tmp
			npre = tmp
		}
	}
	return nroot
}

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/submissions/
func kthLargest(root *TreeNode, k int) int {
	ss := inorderTraversal(root)
	return ss[len(ss)-k]
}