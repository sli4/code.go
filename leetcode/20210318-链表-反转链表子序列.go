package leetcode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head==nil {
		return head
	}
	node := head
	var pre *ListNode
	i:=1
	for ;i<left ;i++ {
		pre = node
		node = node.Next
	}

	var h,t *ListNode
	for ;i<=right ;i++ {
		temp := node
		node = node.Next
		if h==nil {
			h = temp
			t = temp
		} else {
			temp.Next = h
			h = temp
		}
	}
	if pre!=nil {
		pre.Next = h
	} else {
		head = h
	}
	t.Next = node

	return head
}