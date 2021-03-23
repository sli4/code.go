package leetcode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var h,t *ListNode

	for ;l1!=nil && l2!=nil; {
		temp := l1
		if l1.Val > l2.Val {
			temp = l2
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		temp.Next = nil
		if h==nil {
			h = temp
			t = h
		} else {
			t.Next = temp
			t = temp
		}
	}
	x := l1
	if l2!=nil {
		x = l2
	}
	if t==nil {
		t = x
		h = t
	} else {
		t.Next = x
	}
	return h
}