package str

import (
	"math/rand"
)
// https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
func ModifyString(s string) string {
	ns := ""
	for i, c := range s {
		if string(c)=="?" {
			var pre,next int32
			if i-1>=0 {
				pre = int32(ns[i-1])
			}
			if i+1<len(s) {
				next = int32(s[i+1])
			}
			ns = ns+string(GetChar(pre,next))
		} else {
			ns = ns+string(c)
		}
	}
	return ns
}

func GetChar(p,n int32) int32 {
	c := 97 + rand.Int31n(26)
	for ;c==p || c==n;{
		c = 97 + rand.Int31n(26)
	}
	return c
}