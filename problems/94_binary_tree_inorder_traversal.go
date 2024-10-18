package problems

import "leetcode-go/kit"

// https://leetcode.com/problems/binary-tree-inorder-traversal/
type InorderTraversalTreeNode struct {
	Val   int
	Left  *InorderTraversalTreeNode
	Right *InorderTraversalTreeNode
}

func inorderTraversal(root *InorderTraversalTreeNode) []int {
	return inorderTraversalIteratively(root)
}

func inorderTraversalCallStack(root *InorderTraversalTreeNode) []int {
	var result []int
	inorderTraversalCallStackLoop(root, &result)
	return result
}

func inorderTraversalCallStackLoop(n *InorderTraversalTreeNode, inorder *[]int) {
	if n == nil {
		return
	}
	if n.Left != nil {
		inorderTraversalCallStackLoop(n.Left, inorder)
	}
	*inorder = append(*inorder, n.Val)
	if n.Right != nil {
		inorderTraversalCallStackLoop(n.Right, inorder)
	}
}

func inorderTraversalIteratively(root *InorderTraversalTreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := kit.SimpleStack[*InorderTraversalTreeNode]{}
	visited := make(map[*InorderTraversalTreeNode]bool)
	stack.Push(root)
	for !stack.Empty() {
		r, _ := stack.Top()
		if r.Left != nil && !visited[r.Left] {
			stack.Push(r.Left)
			continue
		}
		result = append(result, r.Val)
		visited[r] = true
		stack.Pop()
		if r.Right != nil && !visited[r.Right] {
			stack.Push(r.Right)
			continue
		}
	}
	return result
}
