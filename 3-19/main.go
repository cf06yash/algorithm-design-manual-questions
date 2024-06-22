package main

import (
	"errors"
	"fmt"
	"math"
)

type Node struct {
	left   *Node
	right  *Node
	key    int
	value  string
	height int
}

type AVLTree struct {
	maxElem int
	minElem int
	root    *Node
}

func addNode(root, n *Node) *Node {
	if root == nil {
		return n
	}
	if root.key < n.key {
		root.right = addNode(root.right, n)
	} else if root.key > n.key {
		root.left = addNode(root.left, n)
	} else {
		root.value = n.value
		return root
	}

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	balanceFactor := getBalanceFactor(root)

	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	if balanceFactor < -1 && root.right.key < n.key {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor < -1 && root.right.key > n.key {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor > 1 && root.left.key < n.key {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.key > n.key {
		//LL
		root = rotateRight(root)
	}

	return root
}

func rotateRight(y *Node) *Node {
	x := y.left
	t2 := x.right
	y.left = t2
	x.right = y

	y.height = 1 + max(getHeight(y.right), getHeight(y.left))
	x.height = 1 + max(getHeight(x.right), getHeight(x.left))

	return x
}

func rotateLeft(y *Node) *Node {
	x := y.right
	t2 := x.left
	y.right = t2
	x.left = y

	y.height = 1 + max(getHeight(y.right), getHeight(y.left))
	x.height = 1 + max(getHeight(x.right), getHeight(x.left))

	return x
}

func getBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.left) - getHeight(n.right)
}

func getHeight(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func (t *AVLTree) Add(key int, value string) {
	n := &Node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
	if t.root == nil {
		t.root = n
		t.minElem = key
		t.maxElem = key
		return
	}
	if key > t.maxElem {
		t.maxElem = key
	}
	if key < t.minElem {
		t.minElem = key
	}
	t.root = addNode(t.root, n)
}

func getSuccessor(root *Node) *Node {
	if root == nil {
		return nil
	}
	return findMin(root.right)
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		succ := getSuccessor(root)
		succKey := succ.key
		root.key = succ.key
		root.value = succ.value

		root.right = deleteNode(root.right, succKey)
		return root
	}

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	balanceFactor := getBalanceFactor(root)

	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	if balanceFactor < -1 && getBalanceFactor(root.right) < 0 {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor < -1 && getBalanceFactor(root.right) >= 0 {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor > 1 && getBalanceFactor(root.left) >= 0 {
		//LL
		root = rotateRight(root)
	} else if balanceFactor > 1 && getBalanceFactor(root.left) < 0 {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	}

	return root
}

func findMax(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	} else {
		return findMax(root.right)
	}
}

func findMin(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	} else {
		return findMin(root.left)
	}
}

func (t *AVLTree) Delete(key int) {
	if t.root == nil {
		return
	}
	t.root = deleteNode(t.root, key)
	minNode := findMin(t.root)
	if minNode == nil {
		t.minElem = math.MaxInt
	} else {
		t.minElem = minNode.key
	}
	maxNode := findMax(t.root)
	if maxNode == nil {
		t.maxElem = math.MinInt
	} else {
		t.maxElem = maxNode.key
	}
}

func (t *AVLTree) getMin() (int, error) {
	if t.root == nil {
		return t.minElem, errors.New("empty")
	}
	return t.minElem, nil
}

func (t *AVLTree) getMax() (int, error) {
	if t.root == nil {
		return t.maxElem, errors.New("empty")
	}
	return t.maxElem, nil
}

func (t *AVLTree) printTree() {
	printNodeDetails(t.root)
}

func printNodeDetails(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("Node: %d, height: %d, balance_factor: %d\n", node.key, node.height, getBalanceFactor(node))
	if node.left != nil {
		printNodeDetails(node.left)
	}
	if node.right != nil {
		printNodeDetails(node.right)
	}
}

func printMinMax(t *AVLTree) {
	fmt.Println("-------------------------")
	t.printTree()
	m, ok := t.getMin()
	if ok == nil {
		fmt.Printf("Min Elem %d\n", m)
	} else {
		fmt.Printf("%s\n", ok)
	}
	m, ok = t.getMax()
	if ok == nil {
		fmt.Printf("Max Elem %d\n", m)
	} else {
		fmt.Printf("%s\n", ok)
	}
}

func main() {
	/*
		Suppose you have access to a balanced dictionary data structure that supports each of the operations search, insert, delete, minimum, maximum, successor, and predecessor in O(log n) time. Explain how to modify the insert
		and delete operations so they still take O(log n) but now minimum and maximum take O(1) time. (Hint: think in terms of using the abstract dictionary
		operations, instead of mucking about with pointers and the like.)
	*/
	t := &AVLTree{}
	t.Add(1, "yellow")
	printMinMax(t)
	t.Delete(1)
	t.Add(2, "string")
	t.Add(3, "string")
	printMinMax(t)
	t.Add(4, "yellow")
	printMinMax(t)
	t.Delete(4)
	t.Delete(4)
	t.Delete(2)
	t.Add(6, "lol")
	t.Add(6, "aloha")
	t.Add(7, "kay")
	t.Add(9, "pol")
	t.Add(10, "jop")
	t.Add(1, "ror")
	printMinMax(t)
}
