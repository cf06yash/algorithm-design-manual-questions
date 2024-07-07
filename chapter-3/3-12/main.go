package main

import (
	"fmt"
	"math/rand"
)

type Queue struct {
	items []*Pair
}

type Pair struct {
	n      *Node
	height int
}

func (q *Queue) enqueue(n *Node, h int) {
	q.items = append(q.items, &Pair{
		n:      n,
		height: h,
	})
}

func (q *Queue) dequeue() *Pair {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func insertNode(root, n *Node) *Node {
	if root == nil {
		return n
	}
	randomNumber := rand.Intn(2) + 1
	if randomNumber == 1 {
		root.left = insertNode(root.left, n)
	} else {
		root.right = insertNode(root.right, n)
	}
	return root
}

func (t *Tree) Add(val int) {
	t.root = insertNode(t.root, &Node{
		val:   val,
		left:  nil,
		right: nil,
	})
}

func (t *Tree) getHeight() int {
	var q *Queue = &Queue{
		items: make([]*Pair, 0),
	}
	q.enqueue(t.root, 0)
	var max int = 0
	for len(q.items) > 0 {
		var p *Pair = q.dequeue()
		if p == nil {
			continue
		}
		h := p.height
		if h > max {
			max = h
		}
		n := p.n
		if n.left != nil {
			q.enqueue(n.left, h+1)
		}
		if n.right != nil {
			q.enqueue(n.right, h+1)
		}
	}

	return max
}

func main() {
	/*
		The maximum depth of a binary tree is the number of nodes on the path
		from the root down to the most distant leaf node. Give an O(n) algorithm to
		find the maximum depth of a binary tree with n nodes.
	*/
	var n = 97
	t := &Tree{}
	for i := 1; i <= n; i++ {
		t.Add(i)
	}
	fmt.Printf("Height of the tree is %d\n", t.getHeight())
}
