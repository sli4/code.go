package leetcode




// 交易多次
type Stack struct {
	data []int
	index int
}
func NewStack(n int) *Stack {
	return &Stack{
		data: make([]int, n),
		index: 0,
	}
}
func (s *Stack) In(a int) {
	s.data[s.index] = a
	s.index++
}
func (s *Stack) Out() {
	if s.index>0 {
		s.index--
	}
}

func maxProfitOnce(prices []int) int {
	s := NewStack(len(prices))
	max := 0
	for i, p := range prices {
		if i==0 {
			s.In(p)
			continue
		}
		if s.data[s.index-1] >= p {
			s.In(p)
		} else {
			max = max + p-s.data[s.index-1]
			s.Out()
		}
	}
	return max
}