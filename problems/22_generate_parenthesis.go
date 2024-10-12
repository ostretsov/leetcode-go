package problems

import (
	"strings"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var res []string
	combLen := n * 2
	varCombLen := combLen - 2
	pcMax := 1 << varCombLen
	for pc := range pcMax {
		if !zeroOnesEquals(pcMax, pc) {
			continue
		}
		comb := (pc << 1) | 1
		if !validParenthesis(combLen, comb) {
			continue
		}
		res = append(res, toParenthesis(combLen, comb))
	}
	return res
}

func zeroOnesEquals(pcMax, pc int) bool {
	var zeroCnt, oneCnt byte
	for mask := 1; mask < pcMax; mask = mask << 1 {
		if mask&pc > 0 {
			oneCnt++
			continue
		}
		zeroCnt++
	}
	return zeroCnt == oneCnt
}

// 1 = "(", 0 = ")"
func validParenthesis(combLen, comb int) bool {
	stack := kit.SimpleStack[struct{}]{}
	for i := 0; i < combLen; i++ {
		mask := 1 << i
		if comb&mask > 0 {
			stack.Push(struct{}{})
			continue
		}
		_, ok := stack.Pop()
		if !ok {
			return false
		}
	}
	return stack.Empty()
}

func toParenthesis(combLen, comb int) string {
	var str strings.Builder
	for i := 0; i < combLen; i++ {
		mask := 1 << i
		if comb&mask > 0 {
			str.WriteString("(")
			continue
		}
		str.WriteString(")")
	}
	return str.String()
}
