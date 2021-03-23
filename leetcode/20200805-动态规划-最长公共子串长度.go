package leetcode

import "fmt"

var la = "mitcmu"
var lb = "mtacnu"
var lm = len(a)
var ln = len(b)
var ll = 0
func longestCommonLenth(i,j,length int) {
	if i==lm || j==ln {
		if length > ll {ll=length}
		return
	}
	if la[i]==lb[j] {
		longestCommonLenth(i+1,j+1, length+1)
	} else {
		longestCommonLenth(i+1,j,length )
		longestCommonLenth(i,j+1,length)
	}
}

func LongestCommonLenth(x,y string) {
	la = x
	lb = y
	longestCommonLenth(0,0,0)
	fmt.Println(ll)
}