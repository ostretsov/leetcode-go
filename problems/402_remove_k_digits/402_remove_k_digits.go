package _02_remove_k_digits

import (
	"math"
	"strconv"
	"strings"

	"leetcode-go/kit"
)

func removeKdigits(num string, k int) string {
	return removeKdigitsStack(num, k)
}

func removeKdigitsStack(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	stack := kit.SimpleStack[int]{}
	numr := []rune(num)
	for i, c := range numr {
		if k == 0 {
			break
		}
		for !stack.Empty() && k > 0 {
			ti := stack.MustTop()
			tn := int(num[ti] - '0')
			curn := int(c - '0')
			if curn >= tn {
				break
			}
			numr[ti] = '_'
			stack.Pop()
			k--
		}
		stack.Push(i)
	}
	var result string
	for _, r := range numr {
		if r != '_' {
			result += string(r)
		}
	}
	if k > 0 {
		result = result[:len(result)-k]
	}
	result = strings.TrimLeft(result, "0")
	if result == "" {
		result = "0"
	}
	return result
}

// It was my first attempt and it fails tests
func removeKdigitsV1NaiveAndWrong(num string, k int) string {
	if k == len(num) {
		return "0"
	}

	type numPart struct {
		val     int
		origIdx int
	}
	maxNumRank := len(num) - k
	rankGroup := make(map[int][]numPart)
	for nr := 0; nr < maxNumRank; nr++ {
		for i := nr; i < k+1; i++ {
			n := int(num[i] - '0')
			if n == 0 {
				continue
			}
			rankGroup[maxNumRank-nr-1] = append(rankGroup[maxNumRank-nr-1], numPart{
				val:     int(math.Pow10(maxNumRank-nr-1)) * n,
				origIdx: i,
			})
		}
	}

	deletedIdx := make([]int, len(num))
	for i := maxNumRank - 1; i >= 0 && k > 0; i-- {
		var maxNumPart numPart
		for _, np := range rankGroup[i] {
			if deletedIdx[np.origIdx] == 1 {
				continue
			}
			if np.val > maxNumPart.val {
				maxNumPart = np
			}
		}
		if maxNumPart.val > 0 {
			k--
			deletedIdx[maxNumPart.origIdx] = 1
		}
	}
	var result string
	for i, c := range num {
		if deletedIdx[i] == 1 {
			continue
		}
		result += string(c)
	}
	r, _ := strconv.Atoi(result)
	if r == 0 {
		return "0"
	}
	result = strings.TrimLeft(result, "0")
	return result
}
