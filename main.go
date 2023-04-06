package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (t *Tree) createTree() {
	fmt.Println("Create new tree")
	fmt.Println("----------------")
	reader := bufio.NewReader(os.Stdin)
	key, _ := getInput("Enter a new key for root node: ", reader)
	//key convert into integer
	keyInt, keyErr := strconv.Atoi(key)
	if keyErr != nil {
		fmt.Println("The key must be a number...")
		t.createTree()
	} else {
		value, _ := getInput("Enter a value for root node: ", reader)
		if t.root == nil {
			t.root = &Node{keyInt, value, nil, nil}
		}
		fmt.Printf("Created a tree with root node key is '%v' and its value is '%v' \n", keyInt, value)
		fmt.Println("-----------------------------------------------------------------------")
	}

}
func promptOptions(rn *Node) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option \n i - Insert a value to the tree \n s - Search a  value from the three \n d - Deleate a value from the tree \n e - Exit \n", reader)

	switch opt {
	case "i":
		key, _ := getInput("Enter a key: ", reader)
		keyInt, keyErr := strconv.Atoi(key)
		if keyErr != nil {
			fmt.Println("The key must be a number...")
			promptOptions(rn)
		} else {
			value, _ := getInput("Enter a value for this key: ", reader)
			insertNode(rn, keyInt, value)
			fmt.Printf("Value inserted with key is '%v' and its value is '%v' \n", keyInt, value)
			promptOptions(rn)
		}

	case "s":
		key, _ := getInput("Enter key for searching: ", reader)
		keyInt, keyErr := strconv.Atoi(key)
		if keyErr != nil {
			fmt.Println("The key must be a number...")
			promptOptions(rn)
		} else {
			value := search(rn, keyInt)
			fmt.Println("-------------------------------")
			fmt.Printf("Value for the key '%v' is : %v \n", keyInt, value)
			fmt.Println("-------------------------------")
			promptOptions(rn)
		}
	case "d":
		key, _ := getInput("Enter key for deleting value: ", reader)
		keyInt, keyErr := strconv.Atoi(key)
		if keyErr != nil {
			fmt.Println("The key must be a number...")
			promptOptions(rn)
		} else {
			deleteNode(rn, keyInt)
			fmt.Println("-------------------------------")
			fmt.Println("Deleted Succesfully...")
			fmt.Println("-------------------------------")
			promptOptions(rn)
		}
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(rn)
	}
}
func main() {
	tree := Tree{}
	tree.createTree()
	promptOptions(tree.root)

}
