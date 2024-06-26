package main

import (
	"fmt"
)

type Node struct {
	left      *Node
	right     *Node
	height    int
	remWeight float64
	count     int
}

type AVLTree struct {
	root *Node
}

func createNewNode(weight float64) *Node {
	return &Node{
		remWeight: 1 - weight,
		left:      nil,
		right:     nil,
		height:    1,
		count:     1,
	}
}

func search(root *Node, weight float64) *Node {
	if root == nil {
		return root
	}
	if root.remWeight >= weight && root.left == nil {
		return root
	} else if root.remWeight >= weight {
		n := search(root.left, weight)
		if n != nil {
			return n
		} else {
			return root
		}
	} else {
		return search(root.right, weight)
	}
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

func updateNode(root *Node) {
	if root == nil {
		return
	}
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
}

func addNode(root, n *Node) *Node {
	if root == nil {
		return n
	}
	if root.remWeight < n.remWeight {
		root.right = addNode(root.right, n)
	} else if root.remWeight > n.remWeight {
		root.left = addNode(root.left, n)
	} else {
		root.count++
		return root
	}

	updateNode(root)
	balanceFactor := getBalanceFactor(root)
	if balanceFactor <= 1 && balanceFactor >= 1 {
		return root
	}

	if balanceFactor < -1 && root.right.remWeight < n.remWeight {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor < -1 && root.right.remWeight > n.remWeight {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor > 1 && root.left.remWeight < n.remWeight {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.remWeight > n.remWeight {
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

func getMin(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.left == nil {
		return root
	}
	return getMin(root.left)
}

func deleteNode(root *Node, remWeight float64) *Node {
	if root == nil {
		return root
	}
	if root.remWeight > remWeight {
		root.left = deleteNode(root.left, remWeight)
	} else if root.remWeight < remWeight {
		root.right = deleteNode(root.right, remWeight)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			succ := getMin(root.right)
			sucKey := succ.remWeight
			root.remWeight = succ.remWeight
			root.right = deleteNode(root.right, sucKey)
		}
	}

	updateNode(root)
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

func (t *AVLTree) Add(n *Node) {
	t.root = addNode(t.root, n)
}

func (t *AVLTree) Delete(weight float64) {
	t.root = deleteNode(t.root, weight)
}

func (t *AVLTree) Search(weight float64) *Node {
	return search(t.root, weight)
}

func addItem(old *Node, weight float64) *Node {
	if old == nil {
		return createNewNode(weight)
	}
	return &Node{
		left:      nil,
		right:     nil,
		count:     1,
		height:    1,
		remWeight: old.remWeight - weight,
	}
}

func sumNodeCount(root *Node, count *int) {
	if root == nil {
		return
	}
	(*count) += root.count
	sumNodeCount(root.left, count)
	sumNodeCount(root.right, count)
}

func (t *AVLTree) Count() int {
	count := 0
	sumNodeCount(t.root, &count)
	return count
}

func pack(t *AVLTree, weight float64) {
	var bin *Node = t.Search(weight)
	var updatedBin *Node
	if bin == nil {
		bin = createNewNode(weight)
		updatedBin = bin
	} else {
		if bin.count > 1 {
			bin.count--
			updatedBin = addItem(bin, weight)
		} else {
			updatedBin = addItem(bin, weight)
			t.Delete(bin.remWeight)
		}
	}
	t.Add(updatedBin)
}

func main() {
	/*
		In the bin-packing problem, we are given n objects, each weighing at most
		1 kilogram. Our goal is to find the smallest number of bins that will hold the n
		objects, with each bin holding 1 kilogram at most.
		• The best-fit heuristic for bin packing is as follows. Consider the objects
		in the order in which they are given. For each object, place it into the
		partially filled bin with the smallest amount of extra room after the object is inserted. If no such bin exists, start a new bin. Design an algorithm that implements the best-fit heuristic (taking as input the n weights
		w1, w2, ..., wn and outputting the number of bins used) in O(n log n) time.

		// I am just solving the first part
	*/
	t := &AVLTree{}
	var w []float64 = []float64{0.1, 0.3, 0.6, 0.4, 0.2, 0.9, 0.1, 0.8}
	for _, i := range w {
		pack(t, i)
	}
	fmt.Printf("%d\n", t.Count())
}
