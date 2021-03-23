package leetcode

import "fmt"

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}

func Rob(nums []int) int {
	n := len(nums)
	sum := make([]int,n+3)
	sum[n] = 0
	sum[n+1] = 0
	sum[n+2] = 0
	for i:=n-1;i>=0;i-- {
		sum[i] = nums[i] + max(sum[i+2],sum[i+3])
		fmt.Println("Sum[",i,"]=",sum[i])
	}
	return max(sum[0],sum[1])
}
