package leetcode

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/submissions/

//https://leetcode-cn.com/problems/maximum-subarray/
//https://leetcode-cn.com/problems/contiguous-sequence-lcci/

func maxSubArray(nums []int) int {
	start := 0
	end :=0
	sum := nums[start]
	max := nums[end]
	for i:=1;i<len(nums);i++ {
		if sum<0 {
			start = i
			end = i
			sum = nums[i]
		} else {
			end = i
			sum = sum+nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}