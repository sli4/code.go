package leetcode

import "fmt"

// m = len(a), n = len(b)
// a: mitcmu
// b: mtacnu
var a = "mitcmu"
var b = "mtacnu"
var m = len(a)
var n = len(b)
var minDist = m+n
func lwstBT(i,j,edist int) {
	fmt.Println("m=",m, "(i=",i,"),n=",n, "(j=", j, ")", ",minDist=", minDist)
	if i==m || j==n {// 到达队尾
		if j<n {edist += n-j}  // b未遍历完，删除多余字段
		if i<m {edist += m-i}  // b遍历完，增加多与字段
		if edist < minDist {minDist=edist}
		return
	}

	if a[i]==b[j] {
		lwstBT(i+1,j+1, edist)
	} else {
		lwstBT(i,j+1,edist+1)    // 删除b[j]
		lwstBT(i+1,j, edist+1)   // 在b[j]前加入a[i]
		lwstBT(i+1,j+1, edist+1) // 把b[j]改为a[i]
	}
}

func LwstBT(x,y string) {
	a = x
	b = y

	lwstBT(0,0,0)
	fmt.Println(a, ",", b, ":", minDist)
}