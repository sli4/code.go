package leetcode

import "sort"

// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for first:=0;first<len(nums); first++{
		if first>0 && nums[first]==nums[first-1] {
			continue
		}
		third := len(nums)-1

		for second:=first+1;second<len(nums);second++ {
			if second>first+1 && nums[second]==nums[second-1] {
				continue
			}
			for second<third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[first] + nums[second] + nums[third] == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

// 时间复杂度：O(n^2)