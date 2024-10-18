package _50_eval_rpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	ast.Equal(6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	ast.Equal(22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
