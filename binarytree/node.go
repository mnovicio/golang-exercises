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

type NodeWithLevel struct {
	node  *Node
	level int
}

func (n *Node) GetLevel(val int) int {
	queue := list.New()

	queue.PushBack(&NodeWithLevel{node: n, level: 1})

	for queue.Len() > 0 {
		element := queue.Front()
		nodeWithLevel := queue.Remove(element).(*NodeWithLevel)

		if nodeWithLevel.node.Value == val {
			return nodeWithLevel.level
		}

		if nodeWithLevel.node.Left != nil {
			queue.PushBack(&NodeWithLevel{node: nodeWithLevel.node.Left, level: nodeWithLevel.level + 1})
		}

		if nodeWithLevel.node.Right != nil {
			queue.PushBack(&NodeWithLevel{node: nodeWithLevel.node.Right, level: nodeWithLevel.level + 1})
		}
	}

	return 0
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

func (n *Node) GetHopDFS(val int) int {
	hop, found := n.GetTraversalCountDFS(val, 0)
	if found {
		return hop
	}

	return 0
}

func (n *Node) GetTraversalCountDFS(val, hop int) (int, bool) {
	// fmt.Println("node is now:", n)
	// fmt.Println("hop is now:", hop)

	if n.Value == val {
		return hop + 1, true
	}

	found := false
	if n.Left != nil {
		// fmt.Println("checking left")
		hop, found = n.Left.GetTraversalCountDFS(val, hop+1)
		// hop = ihop
		if found {
			return hop, found
		}
	}

	if n.Right != nil {
		// fmt.Println("checking right")
		hop, found = n.Right.GetTraversalCountDFS(val, hop+1)
		// hop = ihop
		if found {
			return hop, found
		}
	}

	return hop, false
}
