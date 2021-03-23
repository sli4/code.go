package leetcode

import "fmt"

func NumDistinct(s string, t string) int {
	return numDistinctWithIndex(s,t,0,0)
}

func numDistinctWithIndex(s,t string, i,j int) int {
	sum := 0
	fmt.Printf("I:%d, J:%d, S:%s, T:%s\n", i,j, string(s[i]), string(t[j]))
	for k,c := range s {
		if k < i {
			continue
		}
		if c== int32(t[j]) {
			if j==len(t)-1 {
				sum += 1
			} else {
				sum += numDistinctWithIndex(s, t, k+1, j+1)
			}
		}
	}
	return sum
}