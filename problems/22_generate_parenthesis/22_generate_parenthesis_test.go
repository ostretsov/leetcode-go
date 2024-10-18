package _2_generate_parenthesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	ast := assert.New(t)

	got := generateParenthesis(3)
	ast.Len(got, 5)
	ast.Contains(got, "((()))")
	ast.Contains(got, "(()())")
	ast.Contains(got, "(())()")
	ast.Contains(got, "()(())")
	ast.Contains(got, "()()()")

	got = generateParenthesis(1)
	ast.Len(got, 1)
	ast.Contains(got, "()")

	got = generateParenthesis(8)
	ast.Len(got, 1430)
}
