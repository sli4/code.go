package leetcode

func IntegerBreak(n int) int {
	fs := make([]int, n+1)
	fs[0] = 1
	fs[1] = 1

	for i:=2; i<=n; i++ {
		curMax := 1
		for j:=1; j<i;j++ {
			tmp := j*fs[i-j]
			if tmp>curMax {
				curMax = tmp
			}
			tmp = j * (i-j)
			if tmp>curMax {
				curMax = tmp
			}

		}
		fs[i]=curMax
	}
	return fs[n]
}




// f(n)为n拆分数乘积最大值
// f(i) = max{j*f(i-j), j*(i-j)}
//https://leetcode-cn.com/problems/integer-break/solution/zheng-shu-chai-fen-by-leetcode-solution/
