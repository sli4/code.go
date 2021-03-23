package leetcode

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n<= 2 {
		return 0
	}
	f := make([]int, n+1)
	f[n] =0
	f[n-1] = cost[n-1]
	for i := n-2; i>=0; i-- {
		f[i] = cost[i] + min(f[i+1], f[i+2])
	}
	return min(f[0], f[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}