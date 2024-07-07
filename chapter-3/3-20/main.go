package main

import "fmt"

type Node struct {
	left   *Node
	right  *Node
	height int
	key    int
	size   int
}

type AVLTree struct {
	root *Node
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return root.height
}

func getBalanceFactor(root *Node) int {
	if root == nil {
		return 0
	}

	return getHeight(root.left) - getHeight(root.right)
}

func size(root *Node) int {
	if root == nil {
		return 0
	}
	return root.size
}

func updateNode(root *Node) {
	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	root.size = 1 + size(root.left) + size(root.right)
}

func addNode(root, n *Node) *Node {
	if root == nil {
		return n
	}
	if n.key > root.key {
		root.right = addNode(root.right, n)
	} else if n.key < root.key {
		root.left = addNode(root.left, n)
	}
	updateNode(root)
	var balanceFactor = getBalanceFactor(root)
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
	} else if balanceFactor > 1 && root.left.key > n.key {
		//LL
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.key < n.key {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	}

	return root
}

func rotateLeft(root *Node) *Node {
	if root == nil {
		return root
	}
	x := root.right
	t2 := x.left
	root.right = t2
	x.left = root

	updateNode(root)
	updateNode(x)

	return x
}

func rotateRight(root *Node) *Node {
	if root == nil {
		return root
	}

	x := root.left
	t2 := x.right
	root.left = t2
	x.right = root

	updateNode(root)
	updateNode(x)

	return x
}

func (t *AVLTree) Add(key int) {
	var n *Node = &Node{
		left:   nil,
		right:  nil,
		height: 1,
		size:   1,
		key:    key,
	}
	if t.root == nil {
		t.root = n
		return
	}
	t.root = addNode(t.root, n)
}

func getMinElem(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	return getMinElem(root.left)
}

func deleteNodeKth(root *Node, k int) *Node {
	if root == nil {
		return nil
	}
	lsize := size(root.left)

	if k <= lsize {
		root.left = deleteNodeKth(root.left, k)
	} else if k > lsize+1 {
		root.right = deleteNodeKth(root.right, k-lsize-1)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		succ := getMinElem(root.right)
		if succ == nil {
			return nil
		}
		root.key = succ.key
		root.right = deleteNodeKth(root.right, 1)
		return root
	}

	updateNode(root)
	balanceFactor := getBalanceFactor(root)
	if balanceFactor <= 1 && balanceFactor >= -1 {
		return root
	}

	if balanceFactor < -1 && getBalanceFactor(root.right) >= 0 {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor < -1 && getBalanceFactor(root.right) < 0 {
		//RR
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

func (t *AVLTree) DeleteKth(k int) {
	if t.root == nil {
		return
	}
	deleteNodeKth(t.root, k)
}

func print(root *Node) {
	if root == nil {
		return
	}
	print(root.left)
	fmt.Printf("Node: %d, height: %d, balanceFactor: %d, size %d\n", root.key, getHeight(root), getBalanceFactor(root), size(root))
	print(root.right)
}

func (t *AVLTree) Print() {
	print(t.root)
}

func main() {
	/*
		Design a data structure to support the following operations:
			• insert(x,T) – Insert item x into the set T.
			• delete(k,T) – Delete the kth smallest element from T.
			• member(x,T) – Return true iff x ∈ T.
		All operations must take O(log n) time on an n-element set.
	*/

	t := &AVLTree{}
	var arr []int = []int{1, 2, 4, 6, 7, 9, 2, 23, 6, 8, 2, 1, 191, 181, 374}
	for _, i := range arr {
		t.Add(i)
	}
	t.Print()
	fmt.Println("Deleting 1st ------------------------------")
	t.DeleteKth(1)
	t.Print()
	fmt.Println("Deleting 2nd ------------------------------")
	t.DeleteKth(2)
	t.Print()
	fmt.Println("Deleting 7th------------------------------")
	t.DeleteKth(7)
	t.Print()
}
