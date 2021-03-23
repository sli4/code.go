package leetcode

func findLengthOfLCIS(nums []int) int {
	max := 0
	grow := []int{}
	for i, v := range nums {
		if i==0 {
			grow = append(grow, v)
			continue
		}
		if v > grow[len(grow)-1] {
			grow = append(grow, v)
		} else {
			if len(grow) > max {
				max = len(grow)
			}
			grow = []int{v}
		}
	}
	if len(grow) > max {
		max=len(grow)
	}
	return max
}