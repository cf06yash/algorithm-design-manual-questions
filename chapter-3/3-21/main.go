package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

type BST struct {
	root *Node
}

func insertNode(root *Node, n *Node) *Node {
	if root == nil {
		return n
	}
	if root.val < n.val {
		root.right = insertNode(root.right, n)
	} else if root.val > n.val {
		root.left = insertNode(root.left, n)
	}
	return root
}

func (t *BST) Insert(k int) {
	t.root = insertNode(t.root, &Node{
		left:  nil,
		right: nil,
		val:   k,
	})
}

func print(root *Node) {
	if root == nil {
		return
	}
	print(root.left)
	ls := 0
	rs := 0
	if root.left != nil {
		ls = root.left.val
	}
	if root.right != nil {
		rs = root.right.val
	}
	fmt.Printf("Node: %d, left %d, right %d\n", root.val, ls, rs)
	print(root.right)
}

func (t *BST) Print() {
	print(t.root)
	fmt.Println("------------------------------------------------")
}

func mergeNodes(s1 *Node, s2 *Node) *Node {
	if s1 == nil {
		return s2
	}
	if s2 == nil {
		return s1
	}

	curr := s1
	for curr.right != nil {
		curr = curr.right
	}

	curr.right = s2
	return s1
}

func (t1 *BST) Merge(t2 *BST) *BST {
	t1.root = mergeNodes(t1.root, t2.root)
	return t1
}

func main() {
	/*
		A concatenate operation takes two sets S1 and S2, where every key in S1 is
		smaller than any key in S2, and merges them. Give an algorithm to concatenate
		two binary search trees into one binary search tree. The worst-case running
		time should be O(h), where h is the maximal height of the two trees.
	*/
	s1 := &BST{}
	s2 := &BST{}

	var arr1 []int = []int{17, 21, 13, 11, 12}
	for _, val := range arr1 {

		s2.Insert(val)

	}

	var arr2 []int = []int{9, 5, 2, 8, 7, 10}
	for _, val := range arr2 {

		s1.Insert(val)

	}
	s1.Print()
	s2.Print()
	z := s1.Merge(s2)
	z.Print()
}
