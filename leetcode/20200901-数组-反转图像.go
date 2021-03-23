package leetcode


func flipAndInvertImage(A [][]int) [][]int {
	for i:=0; i<len(A); i++ {
		for j:=0; j<=len(A[i])/2; j++ {
			k := len(A[i])-1-j
			if j<=k {
				A[i][j], A[i][k] =  getValue(A[i][k]), getValue(A[i][j])
			}

		}
	}
	return A
}

func getValue(a int) int {
	if a==0 {
		return 1
	}
	return 0
}