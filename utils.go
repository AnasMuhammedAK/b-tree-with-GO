package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	//remove white spaces
	return strings.TrimSpace(input), err
}
func insertNode(node *Node, key int, value string) {
	if key < node.key {
		if node.left == nil {
			node.left = &Node{key, value, nil, nil}
			return
		}
		insertNode(node.left, key, value)
	} else if key > node.key {
		if node.right == nil {
			node.right = &Node{key, value, nil, nil}
			return
		}
		insertNode(node.right, key, value)
	} else {
		node.value = value
	}
}
