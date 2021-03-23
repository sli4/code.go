package leetcode


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	to := 0
	r := new(ListNode)
	root := r
	parent := r
	for ;l1!=nil || l2!=nil; {
		v1 := 0
		if l1!=nil {
			v1 = l1.Val
		}
		v2 := 0
		if l2!=nil {
			v2 = l2.Val
		}
		sum := v1+v2+to
		r.Val = sum%10
		if sum>=10 {
			to = sum/10
		} else {
			to = 0
		}

		if l1!=nil {
			l1 = l1.Next
		}
		if l2!=nil {
			l2 = l2.Next
		}


		r.Next = new(ListNode)
		parent = r
		r = r.Next

	}
	if to!=0 {
		r.Val = to
	} else {
		parent.Next = nil
	}
	return root

}