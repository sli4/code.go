package leetcode

import "fmt"
// https://leetcode-cn.com/circle/discuss/uHREEr/
func SortOddEven(h *ListNode) {
	odd, even := splitListNode(h)
	fmt.Println("拆分结束")


	result := mergeListNode(odd,even)
	for ;result!=nil; result=result.Next {
		fmt.Print(result.Val, " ")
	}
}
func mergeListNode(odd, even *ListNode) *ListNode {
	var h,t *ListNode
	for ;odd!=nil && even!=nil;{
		min := odd
		if odd.Val>even.Val {
			min = even
			even = even.Next
		} else {
			odd = odd.Next
		}
		fmt.Println("Min:", min.Val)
		if h==nil {
			h=min
			t = h
		} else {
			t.Next = min
			t = t.Next
		}
	}
	fmt.Println("H:",h.Val, ", T:", t.Val)
	if odd==nil {
		t.Next = even
	} else {
		t.Next = odd
	}
	return h
}
func splitListNode(h *ListNode) (*ListNode, *ListNode) {
	if h==nil || h.Next==nil {
		return h, nil
	}
	odd,even := h, h.Next
	oh := odd
	var oddFlag=true
	h = h.Next.Next
	even.Next = nil
	for ;h!=nil; {
		fmt.Println("Node:",h.Val)
		if oddFlag {
			odd.Next = h
			odd = odd.Next
			oddFlag = false
			h=h.Next
		} else {
			t := h
			h=h.Next
			if t!=nil {
				t.Next = even
				even = t
			}
			oddFlag = true
		}
	}
	odd.Next =nil

	return oh, even
}