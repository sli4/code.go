package leetcode

import (
	"fmt"
	"math"
)
// https://leetcode-cn.com/problems/print-binary-tree/submissions/


func printTree(root * TreeNode) [][]string {
	height := GetHeight(root)

	rn := int(math.Pow(2, float64(height))-1)
	result :=make([][]string, height)
	for i,_ := range result{
		result[i] = make([]string, rn)
	}


	fill(0, 0, rn-1, root, &result)

	return result
}

func fill(row,start,end int, node *TreeNode, result * [][]string) {
	if node==nil {
		return
	}
	mid := (start+end)/2
	(*result)[row][mid] = fmt.Sprint(node.Val)
	fill(row+1, start, mid-1, node.Left, result)
	fill(row+1, mid+1, end, node.Right, result)
}

func GetHeight(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return 1+max(GetHeight(root.Left), GetHeight(root.Right))
}


