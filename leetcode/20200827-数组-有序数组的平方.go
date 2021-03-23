package leetcode


func sortedSquares(A []int) []int {


	mid := 0
	for ;mid<len(A);{
		if A[mid] <0 {
			mid++
		} else {
			break
		}
	}

	i := mid-1
	j := mid
	B := []int{}
	for ;i>=0 && j<len(A); {
		if -A[i] < A[j] {
			B = append(B, A[i]*A[i])
			i--
		} else {
			B = append(B, A[j]*A[j])
			j++
		}
	}

	if i>=0 {
		for ;i>=0; {
			B = append(B, A[i]*A[i])
			i--
		}
	}

	if j<len(A) {
		for ;j<len(A); {
			B = append(B, A[j]*A[j])
			j++
		}
	}

	return B
}

