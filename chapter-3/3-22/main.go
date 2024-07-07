package main

import "fmt"

type Node struct {
	key    int
	size   int
	left   *Node
	right  *Node
	height int
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

func getSize(root *Node) int {
	if root == nil {
		return 0
	}
	return root.size
}

func updateNode(root *Node) {
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
	root.size = 1 + getSize(root.right) + getSize(root.left)
}

func getBalanceFactor(root *Node) int {
	if root == nil {
		return 0
	}
	return getHeight(root.left) - getHeight(root.right)
}

func addNode(root *Node, key int) *Node {
	if root == nil {
		return &Node{
			key:    key,
			left:   nil,
			right:  nil,
			height: 1,
			size:   1,
		}
	}

	if root.key > key {
		root.left = addNode(root.left, key)
	} else if root.key < key {
		root.right = addNode(root.right, key)
	}

	updateNode(root)
	balanceFactor := getBalanceFactor(root)

	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	if balanceFactor < -1 && root.right.key < key {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor < -1 && root.right.key > key {
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
		//RL
	} else if balanceFactor > 1 && root.left.key > key {
		//LL
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.key < key {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	}

	return root
}

func rotateLeft(y *Node) *Node {
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

func printNode(n *Node) {
	if n != nil {
		left := n.left
		right := n.right
		vl := -77
		vr := -77
		if left != nil {
			vl = left.key
		}
		if right != nil {
			vr = right.key
		}
		fmt.Printf("val: %d, height: %d, balance: %d left child-> %d, right child-> %d\n", n.key, n.height, getBalanceFactor(n), vl, vr)
		if n.left != nil || n.right != nil {
			printNode(n.left)
			printNode(n.right)
		}
	}
}

func findKth(root *Node, k int) int {
	if root == nil {
		return 0
	}
	lsize := getSize(root.left)
	fmt.Printf("lsize: %d, k: %d\n", lsize, k)
	if k < lsize+1 {
		return findKth(root.left, k)
	} else if k > lsize+1 {
		return findKth(root.right, k-lsize-1)
	}

	return root.key
}

func (t *AVLTree) getMedian() int {
	if t.root == nil {
		return -1
	}
	k := getSize(t.root)
	fmt.Printf("k == %d\n", k)
	if k%2 == 0 {
		n1 := findKth(t.root, k/2)
		fmt.Printf("k/2 == %d, n1 = %d\n", k/2, n1)
		n2 := findKth(t.root, (k/2)+1)
		fmt.Printf("k/2 +1 == %d, n2 = %d\n", k/2+1, n2)
		return (n1 + n2) / 2
	} else {
		n1 := findKth(t.root, (k+1)/2)
		fmt.Printf("(k+1)/2 == %d, n1 = %d\n", (k+1)/2, n1)
		return n1
	}
}

func (t *AVLTree) Print() {
	printNode(t.root)
}

func (t *AVLTree) Add(key int) {
	t.root = addNode(t.root, key)
}

func main() {
	/*Design a data structure that supports the following two operations:
	• insert(x) – Insert item x from the data stream to the data structure.
	• median() – Return the median of all elements so far.
	*/
	var arr []int = []int{55, 1551, 56, 12, 66, 13, 567, 1122, 78, 19}
	t := &AVLTree{}
	for _, key := range arr {
		t.Add(key)
	}
	t.Print()
	fmt.Printf("Median: %d\n", t.getMedian())
}
