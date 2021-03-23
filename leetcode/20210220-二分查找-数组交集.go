package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	hm := map[int]int{}
	for _, v := range nums1 {
		hm[v]+=1
	}
	result := []int{}
	for _, v := range nums2 {
		if c,ok:= hm[v]; ok && c>0 {
			result = append(result, v)
			hm[v]-=1
		}
	}
	return result
}
