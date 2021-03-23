package leetcode

func transpose(A [][]int) [][]int {
	B := make([][]int, len(A[0]))
	for i,_ := range B {
		B[i] = make([]int, len(A))
	}

	for i,_ := range B {
		for j,_ := range B[i] {
			B[i][j] = A[j][i]
		}
	}
	return B
}