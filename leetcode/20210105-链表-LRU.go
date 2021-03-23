package leetcode


type LruNode struct {
	Key int
	Val int
	Next *LruNode
}

type lru struct {
	head *LruNode
	count int
	limit int
}

func NewLRU(limit int) *lru {
	return &lru{head: nil, count:0, limit: limit}
}
func (l *lru) Set(key,val int) {
	if l.count == l.limit {
		pre := l.head
		p := l.head
		for ;p.Next!=nil; {
			pre = p
			p = p.Next
		}
		pre.Next = nil
		l.count--
	}
	l.head = &LruNode{
		Key: key,
		Val: val,
		Next:l.head,
	}
	l.count++
}
func (l *lru) Get(key int) int {
	p := l.head
	pre := l.head
	for ;p!=nil; {
		if p.Key==key {
			pre.Next = p.Next
			p.Next = l.head
			l.head = p
			return p.Val
		}
		pre = p
		p = p.Next
	}
	return -1
}
// https://www.nowcoder.com/practice/e3769a5f49894d49b871c09cadd13a61?tpId=117&&tqId=35015&rp=1&ru=/ta/job-code-high&qru=/ta/job-code-high/question-ranking
func LRU(operators [][]int, k int) []int {
	result := []int{}
	lru := NewLRU(k)
	for _, arr := range operators {
		if arr[0]==1 {
			lru.Set(arr[1], arr[2])
		} else {
			result = append(result, lru.Get(arr[1]))
		}
	}
	return result
}