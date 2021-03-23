package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
 	Val int
 	Next *ListNode
 }
func listOfDepth(tree *TreeNode) []*ListNode {
	tns := []*TreeNode{tree}
	lns := []*ListNode{}
	for {
		nts := []*TreeNode{}
		var start, pre *ListNode
		for i, n := range tns {
			if n.Left!=nil {
				nts = append(nts, n.Left)
			}
			if n.Right!=nil {
				nts = append(nts, n.Right)
			}
			tmp := &ListNode{n.Val, nil}
			if i==0 {
				pre = tmp
				start = tmp
			} else {
				pre.Next = tmp
				pre = pre.Next
			}
		}
		lns = append(lns, start)
		if len(nts)==0 {
			break
		}
		tns = nts
	}
	return lns
}

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder(root *TreeNode) [][]int {
	tns := []*TreeNode{root}
	lns := [][]int{}
	if root==nil {
		return lns
	}
	for {
		nts := []*TreeNode{}
		arr := []int{}
		if len(tns)==0 {
			break
		}

		for _, n := range tns {

			if n.Left!=nil {
				nts = append(nts, n.Left)
			}
			if n.Right!=nil {
				nts = append(nts, n.Right)
			}

			arr = append(arr, n.Val)
		}
		lns = append(lns, arr)

		tns = nts
	}
	return lns
}