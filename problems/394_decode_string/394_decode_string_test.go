package _94_decode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeString(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("ababab", decodeString("3[ab]"))
	ast.Equal("abababcabc", decodeString("2[ab]2[abc]"))
	ast.Equal("abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
	ast.Equal("accaccacc", decodeString("3[a2[c]]"))
	ast.Equal("abccdcdcdxyz", decodeString("abc3[cd]xyz"))
}
