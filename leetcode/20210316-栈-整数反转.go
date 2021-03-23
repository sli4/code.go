package leetcode

func reverse(x int) int {
	is := &IntStack{
		vals: []int{},
		neg: false,
	}
	is.Push(x)
	return is.Pop()
}

type IntStack struct {
	vals []int
	neg bool
}
func (is *IntStack) Push(x int) {
	if x<0 {
		is.neg = true
		x = -x
	}
	for ;x!=0; {
		r := x%10
		is.vals = append(is.vals, r)
		x=x/10
	}
}
func (is *IntStack) Pop() int {
	nv := 0
	for i:=len(is.vals)-1;i>=0;i-- {
		nv = 10*nv+is.vals[i]
	}
	if is.neg {
		nv = -nv
	}
	if nv >= 4294967296 || nv < -4294967296 {
		return 0
	}
	return nv
}