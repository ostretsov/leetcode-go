package _94_decode_string

import (
	"strconv"
	"strings"

	"leetcode-go/kit"
)

// https://leetcode.com/problems/decode-string/description/
func decodeString(s string) string {
	var decodedStr string
	var repCntStr, repStr string
	var repCnt int
	stack := kit.SimpleStack[state]{}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			if repCnt > 0 {
				stack.Push(state{repCnt, repStr})
				repCntStr = ""
				repCnt = 0
				repStr = ""
			} else if repStr != "" {
				stack.Push(state{1, repStr})
				repCntStr = ""
				repCnt = 0
				repStr = ""
			}
			repCntStr += string(r)
			continue
		}
		if r == '[' {
			repCnt, _ = strconv.Atoi(repCntStr)
			repCntStr = ""
			continue
		}
		if r >= 'a' && r <= 'z' {
			repStr += string(r)
			continue
		}
		if r == ']' {
			if st, ok := stack.Pop(); ok {
				repStr = st.repStr + strings.Repeat(repStr, repCnt)
				repCntStr = ""
				repCnt = st.repCnt
			} else {
				decodedStr += strings.Repeat(repStr, repCnt)
				repStr = ""
				repCnt = 0
				repCntStr = ""
			}
			continue
		}
	}
	return decodedStr + repStr
}

type state struct {
	repCnt int
	repStr string
}
