package main

import "fmt"

type Node struct {
	count  int
	value  int
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

func getBalanceFactor(root *Node) int {
	if root == nil {
		return 0
	}
	return getHeight(root.left) - getHeight(root.right)
}

func addHelper(root *Node, value int) *Node {
	if root == nil {
		return &Node{
			count:  1,
			value:  value,
			left:   nil,
			right:  nil,
			height: 1,
		}
	}

	if root.value < value {
		root.right = addHelper(root.right, value)
	} else if root.value > value {
		root.left = addHelper(root.left, value)
	} else {
		root.count++
		return root
	}

	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	balanceFactor := getBalanceFactor(root)

	if balanceFactor >= -1 && balanceFactor <= 1 {
		return root
	}

	if balanceFactor < -1 && root.right.value > value {
		//RL
		root.right = rotateRight(root.right)
		root = rotateLeft(root)
	} else if balanceFactor < -1 && root.right.value < value {
		//RR
		root = rotateLeft(root)
	} else if balanceFactor > 1 && root.left.value > value {
		//LL
		root = rotateRight(root)
	} else if balanceFactor > 1 && root.left.value < value {
		//LR
		root.left = rotateLeft(root.left)
		root = rotateRight(root)
	}

	return root
}

func rotateRight(root *Node) *Node {
	x := root.left
	t2 := x.right
	root.left = t2
	x.right = root

	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1

	return x
}

func rotateLeft(root *Node) *Node {
	x := root.right
	t2 := x.left
	root.right = t2
	x.left = root

	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1

	return x
}

func inOrderHelper(res *[]int, root *Node) {
	if root == nil {
		return
	}
	inOrderHelper(res, root.left)
	count := root.count
	for i := 0; i < count; i++ {
		*res = append(*res, root.value)
	}
	inOrderHelper(res, root.right)
}

func (t *AVLTree) Add(i int) {
	t.root = addHelper(t.root, i)
}

func (t *AVLTree) Inorder() []int {
	res := []int{}
	inOrderHelper(&res, t.root)
	return res
}

func main() {
	/*
		Consider a sequence S of n integers with many duplications, such that the
		number of distinct integers in S is O(log n). Give an O(n log log n) worst-case
		time algorithm to sort such sequences.
	*/
	arr := []int{5, 3, 2, 5, 2, 3, 69, 69, 5, 2}
	t := &AVLTree{}
	for i := 0; i < len(arr); i++ {
		t.Add(arr[i])
	}
	res := t.Inorder()
	fmt.Printf("Response : %v\n", res)
}
