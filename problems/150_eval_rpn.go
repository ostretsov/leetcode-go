package problems

import (
	"slices"
	"strconv"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := kit.SimpleStack[int]{}
	for _, token := range tokens {
		if slices.Contains([]string{"+", "-", "*", "/"}, token) {
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(execOp(a, b, token))
			continue
		}
		n, _ := strconv.Atoi(token)
		stack.Push(n)
	}
	result, _ := stack.Top()
	return result
}

func execOp(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("invalid op")
}
