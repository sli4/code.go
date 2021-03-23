package leetcode

// https://leetcode-cn.com/problems/sparse-array-search-lcci/
func findString(words []string, s string) int {
	start,end := 0,len(words)-1
	for ;start<=end; {
		mid := (start+end)/2
		if words[mid]=="" {
			continueFlag :=true
			left := mid-1
			for ;left>=0; {
				if words[left]=="" {
					left = left-1
				} else {
					break
				}
			}
			if left>=0 {
				if s==words[left] {
					return left
				}
				if s<words[left] {
					end = left-1
					continueFlag = false
				}
			}

			if continueFlag {
				right := mid+1
				for ;right<len(words); {
					if words[right]=="" {
						right = right+1
					} else {
						break
					}
				}
				if right < len(words) {
					if s==words[right] {
						return right
					}
					if s>words[right] {
						start = mid+1
						continueFlag = false
					}
				}

			}

			if continueFlag {
				return -1
			}

		} else {
			if s==words[mid] {
				return mid
			}
			if s < words[mid] {
				end = mid-1
			}
			if s > words[mid] {
				start = mid+1
			}
		}
	}
	return -1
}

func findStringWithIndex(words []string, s string, start,end int) int {
	if start<0||end>=len(words) {
		return -1
	}
	mid := (start+end)/2
	if words[mid]=="" {
		l := findStringWithIndex(words, s, start, mid-1)
		r := findStringWithIndex(words, s, mid+1, end)
		if l!=-1 {
			return l
		} else if r!=-1 {
			return r
		} else {
			return -1
		}
	} else {
		if s>words[mid] {
			return findStringWithIndex(words, s, mid+1, end)
		} else if s<words[mid] {
			return findStringWithIndex(words, s, start, mid-1)
		} else {
			return mid
		}
	}
}