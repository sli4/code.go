package leetcode

import "fmt"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	} else {
		return left
	}
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil || p==root || q==root{
		return root
	}
	spMap := map[*TreeNode]*TreeNode{
		root: nil,
	}
	ns := []*TreeNode{root}

	for ;len(ns)>0; {
		ss := []*TreeNode{}
		for _, n := range ns {
			if n.Left!=nil {
				spMap[n.Left] = n
				ss = append(ss, n.Left)
			}
			if n.Right !=nil {
				spMap[n.Right] = n
				ss = append(ss, n.Right)
			}
		}
		ns = ss
	}
	fmt.Printf("%+v\n", spMap)
	pa := listAncestor(p, spMap)
	qa := listAncestor(q, spMap)
	i:=len(pa)-1
	j:=len(qa)-1
	for ;i>=0&&j>=0; {
		if pa[i]==qa[j] {
			i--
			j--
		} else {
			break
		}
	}
	if i<0 && j>=0 && p==qa[j]{
		return p
	}
	if j<0 && i>=0 && q==pa[i] {
		return q
	}

	return pa[i+1]

}

func listAncestor(p *TreeNode, spMap  map[*TreeNode]*TreeNode) []*TreeNode {
	parent := spMap[p]
	fmt.Printf("%d parent is %v\n", p.Val, parent)
	ancestors := []*TreeNode{}
	for ;parent!=nil; {
		ancestors = append(ancestors, parent)
		parent = spMap[parent]
	}
	for _, n := range ancestors {
		fmt.Println(n,":", n.Val)
	}
	fmt.Println("***************")
	return ancestors
}