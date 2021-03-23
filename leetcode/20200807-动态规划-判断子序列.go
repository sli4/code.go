package leetcode
// https://leetcode-cn.com/problems/is-subsequence/submissions/
func isSubsequence(s string, t string) bool {
	i,j := 0,0
	return isEqual(i,j,s,t)

}

func isEqual(i,j int,s,t string) bool {
	if i==len(s) {
		return true
	}
	if j==len(t) {
		return false
	}
	if s[i]==t[j] {
		i+=1
		j+=1
	} else {
		j+=1
	}
	return isEqual(i,j, s,t)
}