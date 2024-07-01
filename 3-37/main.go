package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	key   int
}

type BST struct {
	root *Node
}

func createNewNode(a int) *Node {
	return &Node{
		key:   a,
		left:  nil,
		right: nil,
	}
}

func insert(a int, root *Node) *Node {
	if root == nil {
		return createNewNode(a)
	}
	if a < root.key {
		root.left = insert(a, root.left)
	} else {
		root.right = insert(a, root.right)
	}
	return root
}

func (t *BST) Insert(a int) {
	t.root = insert(a, t.root)
}

func equalsHelper(roota, rootb *Node) bool {
	if roota == nil && rootb == nil {
		return true
	}
	if roota == nil || rootb == nil {
		return false
	}
	if roota.key != rootb.key {
		return false
	}

	return equalsHelper(roota.left, rootb.left) && equalsHelper(roota.right, rootb.right)
}

func (ta *BST) Equals(tb *BST) bool {
	return equalsHelper(ta.root, tb.root)
}

func main() {
	//Write a function to determine whether two binary trees are identical. Identical trees have the same key value at each position and the same structure.
	a := []int{1, 9, 3, 5, 8, 4, 11}
	b := []int{1, 9, 3, 5, 8, 4, 11}
	treeA := &BST{}
	treeB := &BST{}

	for i := 0; i < len(a); i++ {
		treeA.Insert(a[i])
	}
	for i := 0; i < len(b); i++ {
		treeB.Insert(b[i])
	}

	if treeA.Equals(treeB) {
		fmt.Printf("Equal trees :D\n")
	} else {
		fmt.Printf("Inequal trees :C\n")
	}

}
