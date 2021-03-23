package leetcode

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	result := LRU([][]int{
		[]int{1,1,1},
		{1,2,2},
		{1,3,2},
		{2,1},
		{1,4,4},
		{2,2},
	}, 3)
	fmt.Println("Result:", result)
}
