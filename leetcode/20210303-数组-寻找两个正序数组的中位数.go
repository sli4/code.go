package leetcode
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/submissions/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	even := false
	index := (len(nums1)+len(nums2))/2
	if (len(nums1)+len(nums2))%2 == 0 {
		even = true
		index--
	}

	i:=0
	j:=0
	k := -1
	var target float64
	for ;i<len(nums1)&&j<len(nums2); {
		if nums1[i]<nums2[j] {
			target = float64(nums1[i])
			i++
		} else {
			target = float64(nums2[j])
			j++
		}
		k++
		if k==index {
			break
		}

	}

	if k<index {
		if i==len(nums1) {
			j=j+index-k-1
			target = float64(nums2[j])
			j=j+1
		} else {
			i=i+index-k-1
			target = float64(nums1[i])
			i++
		}
	}


	if even {
		return (target+float64(getNext(nums1,nums2,i,j)))/2
	} else {
		return target
	}
}

func getNext(nums1,nums2 []int, i,j int) int {
	if i==len(nums1) {
		return nums2[j]
	}
	if j==len(nums2){
		return nums1[i]
	}
	if nums1[i]<nums2[j] {
		return nums1[i]
	} else {
		return nums2[j]
	}
}