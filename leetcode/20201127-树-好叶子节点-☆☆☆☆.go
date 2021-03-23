package leetcode

import "fmt"

//https://leetcode-cn.com/problems/number-of-good-leaf-nodes-pairs/

func CountPairs(root *TreeNode, distance int) int {
	leafs := LeafNode(root)
	for i:=0; i<len(leafs); i++ {
		fmt.Println(leafs[i])
	}
	count := 0
	for i:=0; i<len(leafs); i++ {
		for j:=i+1; j<len(leafs); j++ {
			fmt.Println("A:", leafs[i], ", B:",leafs[j], ", 距离：", getLeafDistance(leafs[i],leafs[j]) )
			if getLeafDistance(leafs[i],leafs[j]) <= distance {
				count+=1
			}
		}
	}

	fmt.Println("Count:", count)
	return count
}

func getLeafDistance(a,b []int) int {
	i :=0
	for {
		if a[i]!=b[i] {
			break
		}

		i++
	}
	return len(a)-i + len(b)-i
}

func LeafNode(root *TreeNode) [][]int{
	result := [][]int{}
	if root.Left!=nil {
		leftLeafs := LeafNode(root.Left)
		for _, v := range leftLeafs {
			leftResult :=[]int{root.Val}
			leftResult = append(leftResult, v...)
			result = append(result, leftResult)
		}
	}
	if root.Right!=nil {
		rightLeafs := LeafNode(root.Right)
		for _, v :=range rightLeafs {
			rightResult := []int{root.Val}
			rightResult = append(rightResult, v...)
			result = append(result, rightResult)
		}
	}
	if root.Left==nil && root.Right==nil {
		result=[][]int{
			{root.Val},
		}
	}
	return result
}