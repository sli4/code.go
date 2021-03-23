package leetcode

func minArray(numbers []int) int {
	start,end := 0, len(numbers)-1
	for ;start<end; {
		mid := (start+end)/2

		if numbers[mid]<numbers[end] {
			end = mid
		} else if numbers[mid]>numbers[end] {
			start = mid+1
		} else {
			end = end-1
		}
	}
	return numbers[start]
}



