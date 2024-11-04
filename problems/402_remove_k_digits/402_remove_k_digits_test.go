package _02_remove_k_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeKdigits(t *testing.T) {
	assert.Equal(t, "1219", removeKdigits("1432219", 3))
	assert.Equal(t, "200", removeKdigits("10200", 1))
	assert.Equal(t, "0", removeKdigits("10", 2))
	assert.Equal(t, "0", removeKdigits("10", 1))
	assert.Equal(t, "0", removeKdigits("100", 1))
	assert.Equal(t, "0", removeKdigits("9", 1))
	assert.Equal(t, "11", removeKdigits("112", 1))
	assert.Equal(t, "12", removeKdigits("12354", 3))
}
