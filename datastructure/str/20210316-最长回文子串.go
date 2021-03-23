package str


import "fmt"

func LongestPalindrome(s string) string {
	longest := ""
	for x:=0; x<len(s); x++ {
		flag, p := palindrom(s, x-1,x+1)
		fmt.Println("----------------------")

		if flag && len(p)>len(longest) {
			longest = p
		}
		fmt.Println("x:",x,",p:",p,"longest:",longest)
		flag, p = palindrom(s, x, x+1)
		if flag && len(p)>len(longest) {
			longest = p
		}
		fmt.Println("x:",x,",p:",p,"longest:",longest)
	}
	return longest
}

// i...j j+1 ... j+j-i+1
func palindrom(s string, i,j int) (bool, string) {
	for ; i>=0&&j<len(s);  {
		if s[i]==s[j] {
			i--
			j++
		} else {
			break
		}
	}
	if i<0 {
		return true, s[:j]
	}else if j>=len(s) {
		return true, s[i+1:]
	} else {
		return true, s[i+1:j]
	}
}
