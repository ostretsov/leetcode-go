package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		n3 := &InorderTraversalTreeNode{Val: 3}
		n2 := &InorderTraversalTreeNode{Val: 2, Left: n3}
		root := &InorderTraversalTreeNode{Val: 1, Right: n2}

		got := inorderTraversal(root)

		assert.Equal(t, []int{1, 3, 2}, got)
	})

	t.Run("example 2", func(t *testing.T) {
		n9 := &InorderTraversalTreeNode{Val: 9}
		n8 := &InorderTraversalTreeNode{Val: 8, Left: n9}
		n7 := &InorderTraversalTreeNode{Val: 7}
		n6 := &InorderTraversalTreeNode{Val: 6}
		n5 := &InorderTraversalTreeNode{Val: 5, Left: n6, Right: n7}
		n4 := &InorderTraversalTreeNode{Val: 4}
		n3 := &InorderTraversalTreeNode{Val: 3, Right: n8}
		n2 := &InorderTraversalTreeNode{Val: 2, Left: n4, Right: n5}
		root := &InorderTraversalTreeNode{Val: 1, Left: n2, Right: n3}

		got := inorderTraversal(root)

		assert.Equal(t, []int{4, 2, 6, 5, 7, 1, 3, 9, 8}, got)
	})
}
