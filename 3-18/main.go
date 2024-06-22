package main

import (
	"fmt"
)

type Node struct {
	left        *Node
	right       *Node
	val         int
	height      int
	successor   *Node
	predecessor *Node
	parent      *Node
}

type AVLTree struct {
	root *Node
}

func NewTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return node.height
}

func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

func insert(root, node *Node) *Node {
	if root == nil {
		return node
	}
	if root.val < node.val {
		if root.right == nil {
			root.right = node
			node.parent = root
			node.predecessor = root
			if root.successor != nil {
				root.successor.predecessor = node
				node.successor = root.successor
			}
			root.successor = node
		} else {
			root.right = insert(root.right, node)
		}
	} else if root.val > node.val {
		if root.left == nil {
			root.left = node
			node.parent = root

			node.successor = root
			if root.predecessor != nil {
				node.predecessor = root.predecessor
				root.predecessor.successor = node
			}
			root.predecessor = node
		} else {
			root.left = insert(root.left, node)
		}
	}
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
	var balanceFactor int = getBalanceFactor(root)

	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	root = balanceTree(root, balanceFactor, node.val)
	return root
}

func balanceTree(root *Node, balanceFactor, key int) *Node {
	if balanceFactor < -1 && root.right.val > key {
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
		//RL
	} else if balanceFactor < -1 && root.right.val < key {
		root = rotateLeft(root)
		//RR
	} else if balanceFactor > 1 && root.left.val < key {
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
		//LR
	} else if balanceFactor > 1 && root.left.val > key {
		root = rotateRight(root)
		// LL
	}

	return root
}

func rotateRight(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2

	x.parent = y.parent
	y.parent = x
	if t2 != nil {
		t2.parent = y
	}

	y.height = 1 + max(getHeight(y.right), getHeight(y.left))
	x.height = 1 + max(getHeight(x.left), getHeight(x.right))
	return x
}

func rotateLeft(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.right
	t2 := x.left
	x.left = y
	y.right = t2

	x.parent = y.parent
	y.parent = x
	if t2 != nil {
		t2.parent = y
	}

	y.height = 1 + max(getHeight(y.left), getHeight(y.right))
	x.height = 1 + max(getHeight(x.left), getHeight(x.right))

	return x
}

func (t *AVLTree) Insert(key int) {
	n := &Node{
		left:   nil,
		right:  nil,
		val:    key,
		height: 0,
	}
	if t.root == nil {
		t.root = n
		return
	}
	t.root = insert(t.root, n)
}

func minElement(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.left == nil {
		return root
	}
	return minElement(root.left)
}

func (t *AVLTree) getMin() *Node {
	return minElement(t.root)
}

func (t *AVLTree) printTree() {
	printNodeDetails(t.root)
}

func printNodeDetails(node *Node) {
	if node == nil {
		return
	}
	sval := 0
	pval := 0
	if node.successor != nil {
		sval = node.successor.val
	}
	if node.predecessor != nil {
		pval = node.predecessor.val
	}
	fmt.Printf("Node: %d, height: %d, pred: %d, succ: %d, balance_factor: %d\n", node.val, node.height, sval, pval, getBalanceFactor(node))
	if node.left != nil {
		printNodeDetails(node.left)
	}
	if node.right != nil {
		printNodeDetails(node.right)
	}
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	if key < root.val {
		root.left = deleteNode(root.left, key)
	} else if key > root.val {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil {
			temp := root.right
			if root.predecessor != nil {
				root.predecessor.successor = root.successor
			}
			if root.successor != nil {
				root.successor.predecessor = root.predecessor
			}
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left

			if root.predecessor != nil {
				root.predecessor.successor = root.successor
			}
			if root.successor != nil {
				root.successor.predecessor = root.predecessor
			}
			root = nil
			return temp
		}

		temp := minElement(root.right)

		root.val = temp.val

		root.right = deleteNode(root.right, temp.val)
	}

	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
	balance := getBalanceFactor(root)

	if balance > 1 && getBalanceFactor(root.left) >= 0 {
		return rotateRight(root)
	}
	if balance > 1 && getBalanceFactor(root.left) < 0 {
		root.left = rotateLeft(root.left)
		return rotateRight(root)
	}

	if balance < -1 && getBalanceFactor(root.right) <= 0 {
		return rotateLeft(root)
	}

	if balance < -1 && getBalanceFactor(root.right) > 0 {
		root.right = rotateRight(root.right)
		return rotateLeft(root)
	}

	return root
}

func (t *AVLTree) delete(key int) {
	deleteNode(t.root, key)
}

func main() {
	/*
		Describe how to modify any balanced tree data structure such that search,
		insert, delete, minimum, and maximum still take O(log n) time each, but successor and predecessor now take O(1) time each. Which operations have to be
		modified to support this?
	*/
	var arr = []int{10, 20, 30, 40, 50, 25}
	t := NewTree()
	for _, element := range arr {
		t.Insert(element)
	}

	fmt.Println("AVL Tree Structure:")
	t.printTree()
	t.delete(30)
	fmt.Println("After deleting value")
	t.printTree()
	m := t.getMin()
	fmt.Println("\nIn-Order Traversal (Successor):")
	for m != nil {
		fmt.Printf("%d ", m.val)
		if m.successor == nil {
			break
		}
		m = m.successor
	}
	fmt.Println()

	fmt.Println("Reverse In-Order Traversal (Predecessor):")
	for m != nil {
		fmt.Printf("%d ", m.val)
		if m.predecessor == nil {
			break
		}
		m = m.predecessor
	}
	fmt.Println()
}
