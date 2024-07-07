package main

import (
	"fmt"
	"strings"
)

type Node struct {
	left   *Node
	right  *Node
	key    string
	height int
}

type AVLTree struct {
	root *Node
}

func addNode(root *Node, key string) *Node {
	if root == nil {
		return &Node{
			left:   nil,
			right:  nil,
			key:    key,
			height: 1,
		}
	}
	if strings.Compare(key, root.key) == 1 {
		root.right = addNode(root.right, key)
	} else if strings.Compare(key, root.key) == -1 {
		root.left = addNode(root.left, key)
	}

	updateNode(root)
	balanceFactor := getBalanceFactor(root)
	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	if balanceFactor < -1 && strings.Compare(key, root.right.key) == -1 {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor < -1 && strings.Compare(key, root.right.key) == 1 {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor > 1 && strings.Compare(key, root.left.key) == 1 {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	} else if balanceFactor > 1 && strings.Compare(key, root.left.key) == -1 {
		//LL
		root = rotateRight(root)
	}

	return root
}

func rotateRight(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.left
	t2 := x.right
	y.left = t2
	x.right = y

	updateNode(y)
	updateNode(x)
	return x
}

func rotateLeft(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.right
	t2 := x.left
	y.right = t2
	x.left = y

	updateNode(y)
	updateNode(x)
	return x
}

func getBalanceFactor(root *Node) int {
	if root == nil {
		return 0
	}
	return getHeight(root.left) - getHeight(root.right)
}

func updateNode(root *Node) {
	if root == nil {
		return
	}
	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return root.height
}

func (t *AVLTree) Add(key string) {
	t.root = addNode(t.root, key)
}

func findPrefix(root *Node, prefix string, count *int, result *[]string) {
	if *count <= 0 {
		return
	}
	if root == nil {
		return
	}
	fmt.Printf("Visted: %s\n", root.key)
	if strings.Compare(root.key[:len(prefix)], prefix) == -1 {
		findPrefix(root.right, prefix, count, result)
	} else if strings.Compare(root.key[:len(prefix)], prefix) == 1 {
		findPrefix(root.left, prefix, count, result)
	} else {
		*result = append(*result, root.key)
		*count = *count - 1
		findPrefix(root.right, prefix, count, result)
		findPrefix(root.left, prefix, count, result)
	}
}

func (t *AVLTree) FindWordsWithPrefix(prefix string, count int) []string {
	var result []string = make([]string, 0)
	findPrefix(t.root, prefix, &count, &result)
	return result
}

func main() {
	/*
	   Assume we are given a standard dictionary (balanced binary search tree)
	   defined on a set of n strings, each of length at most l. We seek to print out all
	   strings beginning with a particular prefix p. Show how to do this in O(ml log n)
	   time, where m is the number of strings
	*/

	t := &AVLTree{}
	var arr []string = []string{"alpha", "satin", "potato", "network", "neuron", "pottassium", "potats", "rabbit", "potu"}
	for _, x := range arr {
		t.Add(x)
	}
	prefixToSearchFor := "pot"
	numberOfResultsNeeded := 1

	results := t.FindWordsWithPrefix(prefixToSearchFor, numberOfResultsNeeded)
	fmt.Printf("%v\n", results)
}
