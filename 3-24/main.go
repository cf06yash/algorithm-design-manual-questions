package main

import (
	"errors"
	"fmt"
)

type Node struct {
	left   *Node
	right  *Node
	key    int
	height int
	size   int
}

func getSize(root *Node) int {
	if root == nil {
		return 0
	}
	return root.size
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
	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	root.size = 1 + getSize(root.left) + getSize(root.right)
}

func createNewNode(key int) *Node {
	return &Node{
		left:   nil,
		right:  nil,
		height: 1,
		size:   1,
		key:    key,
	}
}

func addNode(root *Node, key int) (*Node, error) {
	if root == nil {
		return createNewNode(key), nil
	}
	if root.key > key {
		n, err := addNode(root.left, key)
		if err != nil {
			return nil, err
		}
		root.left = n
	} else if root.key < key {
		n, err := addNode(root.right, key)
		if err != nil {
			return nil, err
		}
		root.right = n
	} else {
		return nil, errors.New("collision")
	}

	updateNode(root)
	balanceFactor := getBalanceFactor(root)
	if balanceFactor <= 1 && balanceFactor >= -1 {
		return root, nil
	}

	if balanceFactor > 1 && root.left.key > key {
		//LL
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.key < key {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	} else if balanceFactor < -1 && root.right.key < key {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor < -1 && root.right.key > key {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	}

	return root, nil
}

func rotateLeft(root *Node) *Node {
	if root == nil {
		return root
	}

	x := root.right
	t2 := x.left
	x.left = root
	root.right = t2
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
	x.right = root
	root.left = t2
	updateNode(root)
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

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	if root.key > key {
		root.left = deleteNode(root.left, key)
	} else if root.key < key {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			succ := getMin(root.right)
			sucKey := succ.key
			root.key = succ.key
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

func (t *AVLTree) Add(key int) error {
	root, err := addNode(t.root, key)
	if err != nil {
		return err
	} else {
		t.root = root
		return nil
	}
}

func (t *AVLTree) Delete(key int) {
	t.root = deleteNode(t.root, key)
}

func (t *AVLTree) Size() int {
	return getSize(t.root)
}

func checkKUniqueDuplicate(arr *[]int, k int) bool {
	if k <= 1 {
		return true
	}
	t := &AVLTree{}
	i := 0
	j := 0

	for ; i < len(*arr); i++ {
		if t.Size() == k {
			t.Delete((*arr)[j])
			j++
		}
		err := t.Add((*arr)[i])
		if err != nil {
			return false
		}
	}
	return true
}

type AVLTree struct {
	root *Node
}

func main() {
	/*
		An array A is called k-unique if it does not contain a pair of duplicate
		elements within k positions of each other, that is, there is no i and j such that
		A[i] = A[j] and |j − i| ≤ k. Design a worst-case O(n log k) algorithm to test if
		A is k-unique.
	*/

	var arr []int = []int{1, 2, 3, 9, 4, 5, 17, 299, 1819, 1, 5}
	k := 5
	fmt.Printf("is array %v %d-unique %t\n", arr, k, checkKUniqueDuplicate(&arr, k))
}
