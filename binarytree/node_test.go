package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintBFS(t *testing.T) {
	tree := &Node{
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

	buf := bytes.NewBufferString("")
	tree.printBFS(buf)

	assert.EqualValues(t, "5 8 3 6 2 4 1 7 9 ", buf.String())
}

func Test_PrintDFS(t *testing.T) {
	tree := &Node{
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

	buf := bytes.NewBufferString("")
	tree.printDFS(buf)

	assert.EqualValues(t, "5 8 6 2 7 9 3 4 1 ", buf.String())
}
