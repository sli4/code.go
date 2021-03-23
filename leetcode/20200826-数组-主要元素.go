package leetcode

// https://leetcode-cn.com/problems/find-majority-element-lcci/


func majorityElement(nums []int) int {
	r := 1
	e := nums[0]
	for i, v := range nums {
		if i==0 {
			continue
		}
		if v==e {
			r++
		} else {
			if r<=0 {
				r=1
				e=v
			} else {
				r--
			}
		}
	}
	if r>0 {
		return e
	}
	return -1
}