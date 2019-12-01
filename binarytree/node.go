package main

import (
	"container/list"
	"fmt"
	"io"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// func (n *Node) AddLeft(val int) {
// 	n.Left = &Node{
// 		Value: val,
// 	}
// }

// func (n *Node) AddRight(val int) {
// 	n.Right = &Node{
// 		Value: val,
// 	}
// }

func (n *Node) printBFS(w io.Writer) {
	queue := list.New()

	queue.PushBack(n)

	// first := true

	for queue.Len() > 0 {
		element := queue.Front()
		node := queue.Remove(element).(*Node)
		// if !first {
		// 	fmt.Fprint(w, " ")
		// } else {
		// 	first = false
		// }
		fmt.Fprintf(w, "%d ", node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

func (n *Node) printDFS(w io.Writer) {
	fmt.Fprintf(w, "%d ", n.Value)

	if n.Left != nil {
		n.Left.printDFS(w)
	}

	if n.Right != nil {
		n.Right.printDFS(w)
	}
}
