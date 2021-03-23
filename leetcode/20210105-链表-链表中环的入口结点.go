package leetcode


// https://www.nowcoder.com/practice/6e630519bf86480296d0f1c868d425ad?tpId=117
// 核心思想：快慢指针+数学计算
func detectCycle( head *ListNode) *ListNode {
	// write code here
	fast := head
	slow := head
	for ;fast!=nil && fast.Next!=nil;{
		fast = fast.Next.Next
		slow = slow.Next
		if fast==slow {
			// 有环

			fast = head
			for ;fast!=slow; {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}