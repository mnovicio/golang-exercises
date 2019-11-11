package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPallindrome(t *testing.T) {
	assert.True(t, isPallindrome("anna"))
	assert.True(t, isPallindrome("annna"))
	assert.False(t, isPallindrome("notanna"))
	assert.False(t, isPallindrome(""))
	assert.True(t, isPallindrome("n"))
}
