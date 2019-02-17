package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := TreeNode{Val: preorder[0], Left: nil, Right: nil}
	loc := 0
	for k, v := range inorder {
		if v == root.Val {
			loc = k
			break
		}
	}
	left, right := inorder[0:loc], inorder[loc+1:]
	root.Left = buildTree(preorder[1:len(left)+1], left)
	root.Right = buildTree(preorder[len(left)+1:], right)
	return &root
}
