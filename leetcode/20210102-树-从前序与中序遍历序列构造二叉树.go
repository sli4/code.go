package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BestBuildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder)==0 || len(inorder)==0{
		return nil
	}
	root:=&TreeNode{Val: preorder[0]}
	for i:=0;i<len(inorder);i++{
		if 	inorder[i]==root.Val{
			root.Left= buildTree(preorder[1:i+1],inorder[:i])
			root.Right= buildTree(preorder[i+1:len(preorder)],inorder[i+1:])
			break
		}
	}
	return root
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	getSonNode(preorder,inorder, root, true,0,len(preorder)-1, 0, len(inorder)-1)
	return root
}

func getSonNode(preorder []int, inorder []int, parent *TreeNode, leftFlag bool, preStart,preEnd, inStart,inEnd int) {
	if preStart>preEnd || inStart>inEnd{
		return
	}
	son := &TreeNode{
		Val: preorder[preStart],
	}

	if parent.Val==son.Val {
		son = parent
	} else {
		if leftFlag {
			parent.Left = son
		} else {
			parent.Right = son
		}
	}
	if preStart==preEnd {
		return
	}

	inMid := inStart
	for ;inMid<=inEnd;inMid++ {
		if inorder[inMid]==son.Val {
			break
		}
	}
	leftSonLen := inMid-inStart

	getSonNode(preorder,inorder, son, true, preStart+1, preStart+leftSonLen, inStart, inStart+leftSonLen-1)
	getSonNode(preorder,inorder, son, false, preStart+leftSonLen+1, preEnd, inMid+1,inEnd)


}