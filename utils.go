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

// InOrder print node of tree
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%d ", root.key)
	InOrder(root.right)
}

// PreOrder print node of tree
func PreOrder(root *Node) {
	if nil == root {
		return
	}
	PreOrder(root.left)
	PreOrder(root.right)
	fmt.Printf("%d ", root.key)
}

// PostOrder print node of tree
func PostOrder(root *Node) {
	if nil == root {
		return
	}
	fmt.Printf("%d ", root.key)
	PostOrder(root.left)
	PostOrder(root.right)
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
func search(rn *Node, key int) string {
	node := rn
	for node != nil {
		if key == node.key {
			return node.value
		} else if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return "Data not found"
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		// case 1: leaf node
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			// case 2: one child
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// case 3: two children
			successor := minkeyNode(root.right)
			root.key = successor.key
			root.right = deleteNode(root.right, successor.key)
		}
	}
	return root
}

func minkeyNode(root *Node) *Node {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current
}
