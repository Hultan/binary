package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {

	// Creating AVL tree and
	// inserting data in it
	root := NewTree(33)
	root = root.Insert(13)
	root = root.Insert(53)
	root = root.Insert(9)
	root = root.Insert(21)
	root = root.Insert(61)
	root = root.Insert(8)
	root = root.Insert(11)

	// Printing AVL Tree
	root.Print("", true)

	// Deleting a node from AVL Tree
	root = root.Delete(13)
	fmt.Println("After deleting ")
	root.Print("", true)
}
