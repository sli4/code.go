package leetcode

func missingNumber(nums []int) int {
	start := 0
	end := len(nums)-1
	mid := (start+end)/2
	for ;start<end; {

		if nums[mid] > mid {
			end = mid
		} else {
			start = mid+1
		}
		mid = (start+end)/2
	}
	if mid==nums[mid] {
		return mid+1
	}
	return mid
}
