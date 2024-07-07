package main

import "fmt"

type Node struct {
	height        int
	balanceFactor int
	left          *Node
	right         *Node
	val           int
}

type AvlTree struct {
	root *Node
}

func insertNode(root *Node, n *Node) *Node {
	if root == nil {
		return n
	} else if n.val > root.val {
		root.right = insertNode(root.right, n)
	} else if n.val < root.val {
		root.left = insertNode(root.left, n)
	} else {
		return root
	}
	updateNode(root)
	root = balance(root, n)
	return root
}

func balance(root, n *Node) *Node {
	if root.balanceFactor <= 1 && root.balanceFactor >= -1 {
		return root
	}
	if root.balanceFactor < -1 && root.right.val > n.val {
		//RL
		root.right = rightRotate(root.right)
		root = leftRotate(root)
	} else if root.balanceFactor < -1 && root.right.val < n.val {
		//RR
		root = leftRotate(root)
	} else if root.balanceFactor > 1 && root.left.val > n.val {
		//LL
		root = rightRotate(root)
	} else if root.balanceFactor > 1 && root.left.val < n.val {
		//LR
		root.left = leftRotate(root.left)
		root = rightRotate(root)
	}

	return root
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func rightRotate(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2

	updateNode(y)
	updateNode(x)
	return x
}

func leftRotate(y *Node) *Node {
	if y == nil {
		return y
	}
	x := y.right
	t2 := x.left
	x.left = y
	y.right = t2

	updateNode(y)
	updateNode(x)
	return x
}

func updateNode(root *Node) {
	root.height = max(getHeight(root.right), getHeight(root.left)) + 1
	root.balanceFactor = getHeight(root.left) - getHeight(root.right)
}

func printNode(n *Node) {
	if n != nil {
		left := n.left
		right := n.right
		vl := -77
		vr := -77
		if left != nil {
			vl = left.val
		}
		if right != nil {
			vr = right.val
		}
		fmt.Printf("val: %d, height: %d, balance: %d left child-> %d, right child-> %d\n", n.val, n.height, n.balanceFactor, vl, vr)
		if n.left != nil || n.right != nil {
			printNode(n.left)
			printNode(n.right)
		}
	}
}

func (t *AvlTree) Print() {
	printNode(t.root)
}

func (t *AvlTree) Insert(val int) {
	var n *Node = &Node{
		height:        1,
		balanceFactor: 0,
		left:          nil,
		right:         nil,
		val:           val,
	}
	if t.root == nil {
		t.root = n
	} else {
		t.root = insertNode(t.root, n)
	}
}

func main() {
	/*
		Describe an O(n)-time algorithm that takes an n-node binary search tree
		and constructs an equivalent height-balanced binary search tree. In a heightbalanced binary search tree, the difference between the height of the left and
		right subtrees of every node is never more than 1.
	*/

	var vals []int = []int{2, 4, 6, 31, 88, 9, 11, 1, 99, 101}
	// just creating an AVL tree out of these values, because I think doing a post order traversal and calling balance and update node function on all nodes as I visit them will solve this
	var t *AvlTree = &AvlTree{}
	for _, i := range vals {
		t.Insert(i)
	}
	t.Print()
}
