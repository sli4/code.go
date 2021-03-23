package leetcode
// https://leetcode-cn.com/problems/range-sum-query-immutable/submissions/
type NumArray struct {
	nums []int
	sums []int
}


func Constructor(nums []int) NumArray {
	sums := make([]int,len(nums))
	for i, x := range nums {
		if i==0 {
			sums[i] = x
		} else {
			sums[i] = sums[i-1] + x
		}
	}

	return NumArray{nums:nums, sums:sums}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i==0 {
		return this.sums[j]
	}
	return this.SumRange(0,j)-this.SumRange(0,i-1)
}

