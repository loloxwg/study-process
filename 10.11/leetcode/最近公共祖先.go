package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root == p || root==q {
		return root
	}
	left:=lowestCommonAncestor2(root.Left,p,q)
	right:=lowestCommonAncestor2(root.Right,p,q)
	if left != nil && right != nil {
		return root
	}
	if left==nil {
		return right
	}
	return left

}


