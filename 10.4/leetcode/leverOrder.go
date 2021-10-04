package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/*
广度优先搜索，层次遍历
*/


func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		n := len(stack)
		var v []int
		for i := 0;i < n;i ++ {
			node := stack[0] //出队列的那个节点
			stack = stack[1:]//出队列
			if node != nil {
				v = append(v, node.Val)
				if node.Left != nil {
					stack = append(stack, node.Left)
				}
				if node.Right != nil {
					stack = append(stack, node.Right)
				}
			}
		}
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
}













func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	var stack []*TreeNode
	stack=append(stack,root)
	for len(stack)!=0{
		n:=len(stack)
		var v []int
		for i := 0; i < n; i++ {
			node:=stack[0]
			stack=stack[1:]
			if node!=nil {
				v=append(v,node.Val)
				if node.Left!=nil {
					stack=append(stack,node.Left)
				}
				if node.Right!=nil {
					stack=append(stack,node.Right)
				}
			}

		}
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
}
















