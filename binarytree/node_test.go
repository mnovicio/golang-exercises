package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &Node{
	Value: 5,
	Left: &Node{
		Value: 8,
		Left: &Node{
			Value: 6,
		},
		Right: &Node{
			Value: 2,
			Left: &Node{
				Value: 7,
			},
			Right: &Node{
				Value: 9,
			},
		},
	},
	Right: &Node{
		Value: 3,
		Left: &Node{
			Value: 4,
		},
		Right: &Node{
			Value: 1,
		},
	},
}

func Test_PrintBFS(t *testing.T) {

	buf := bytes.NewBufferString("")
	tree.printBFS(buf)

	assert.EqualValues(t, "5 8 3 6 2 4 1 7 9 ", buf.String())
}

func Test_GetLevel(t *testing.T) {
	actual := tree.GetLevel(4)
	assert.Equal(t, 3, actual)

	actual = tree.GetLevel(7)
	assert.Equal(t, 4, actual)

	actual = tree.GetLevel(10)
	assert.Equal(t, 0, actual)
}

func Test_PrintDFS(t *testing.T) {
	buf := bytes.NewBufferString("")
	tree.printDFS(buf)

	assert.EqualValues(t, "5 8 6 2 7 9 3 4 1 ", buf.String())
}

func Test_GetHopDFS(t *testing.T) {
	actual := tree.GetHopDFS(5)
	assert.Equal(t, 1, actual)

	actual = tree.GetHopDFS(7)
	assert.Equal(t, 5, actual)

	actual = tree.GetHopDFS(2)
	assert.Equal(t, 4, actual)

	actual = tree.GetHopDFS(10)
	assert.Equal(t, 0, actual)
}
