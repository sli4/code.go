package leetcode



func InsertWithoutRecurse(node *TreeNode, val int) *TreeNode {
	if node==nil {
		node = &TreeNode{Val: val}
		return node
	}
	root := node
	for {
		if node.Val > val {
			if node.Left==nil {
				node.Left = &TreeNode{Val: val}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right==nil {
				node.Right = &TreeNode{Val: val}
				break
			} else {
				node = node.Right
			}
		}
	}
	return root
}

func (n *TreeNode) Insert( val int) *TreeNode {
	if n==nil {
		n = &TreeNode{Val: val}
		return n
	}
	if n.Val>val {
		if n.Left==nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right==nil {
			n.Right= &TreeNode{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
	return n
}
func InorderTraversal(root *TreeNode) []int {
	if root==nil {
		return []int{}
	}
	result := InorderTraversal(root.Left)
	result = append(result, root.Val)
	result =append(result, InorderTraversal(root.Right)...)
	return result
}


func Delete(root *TreeNode, val int) {
	node := root
	var parent *TreeNode

	for {
		if node==nil {
			break
		}
		if val < node.Val {
			parent = node
			node = node.Left
		}
		if val > node.Val {
			parent = node
			node = node.Right
		}
		if node.Val==val {
			break
		}
	}

	if node.Left!=nil && node.Right!=nil {
		minPP := node
		minP := node.Right
		for {
			if minP.Left==nil {
				break
			}
			minPP = minP
			minP = minP.Left
		}

		node.Val = minP.Val
		if minPP==node {
			node.Right=minP.Right
		} else {
			minPP.Left = minP.Right
		}
	}
	//
	var child *TreeNode
	if node.Left==nil && node.Right==nil { // 叶子节点
		child = nil
	}else if node.Left==nil {  // 左子树为空
		child = node.Right
	} else if node.Right==nil { // 右子树为空
		child = node.Left
	}
	if parent==nil {
		root = child
	} else {
		if parent.Left==node {
			parent.Left = child
		}else{
			parent.Right = child
		}
	}
}

func popMinLeftSon(parent,node *TreeNode) *TreeNode {

	for {
		if node.Left==nil {
			break
		}
		parent = node
		node = node.Left
	}
	parent.Left = nil
	return node
}