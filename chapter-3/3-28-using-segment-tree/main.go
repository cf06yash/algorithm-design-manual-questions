package main

import "fmt"

type SegmentTree struct {
	tree     []int
	elements int
}

func getLeftChildPosition(pos int) int {
	return pos + 1
}

func getRightChildPosition(l, mid, pos int) int {
	nodesLeftTree := 2*(mid-l+1) - 1
	return pos + nodesLeftTree + 1
}

func build(tree *[]int, l, r, pos int, arr []int) int {
	if l == r {
		(*tree)[pos] = arr[l]
		return (*tree)[pos]
	}
	mid := (l + r) / 2
	left := getLeftChildPosition(pos)
	right := getRightChildPosition(l, mid, pos)
	(*tree)[pos] = build(tree, l, mid, left, arr) + build(tree, mid+1, r, right, arr)
	return (*tree)[pos]
}

func CreateTree(arr []int) *SegmentTree {
	t := &SegmentTree{
		tree:     make([]int, 2*len(arr)),
		elements: len(arr),
	}
	build(&t.tree, 0, len(arr)-1, 0, arr)
	return t
}

func (t *SegmentTree) Update(idx, value, pos, l, r int) {
	if idx >= t.elements {
		return
	}
	if l == r {
		t.tree[pos] += value
		return
	}
	mid := (l + r) / 2
	if idx <= mid {
		t.Update(idx, value, getLeftChildPosition(pos), l, mid)
	} else {
		t.Update(idx, value, getRightChildPosition(l, mid, pos), mid+1, r)
	}
	t.tree[pos] = t.tree[getLeftChildPosition(pos)] + t.tree[getRightChildPosition(l, mid, pos)]
}

func (t *SegmentTree) Query(ql, qr, l, r, pos int) int {
	if ql <= l && r <= qr {
		return t.tree[pos]
	}
	if qr < l || ql > r {
		return 0
	}
	mid := (l + r) / 2
	return t.Query(ql, qr, l, mid, getLeftChildPosition(pos)) + t.Query(ql, qr, mid+1, r, getRightChildPosition(l, mid, pos))
}

func main() {
	/*
		Let A[1..n] be an array of real numbers. Design an algorithm to perform
		any sequence of the following operations:

		• Add(i,y) – Add the value y to the ith number.
		• Partial-sum(i) – Return the sum of the first i numbers, that is, i
		j=1 A[j].
		There are no insertions or deletions; the only change is to the values of the numbers. Each operation should take O(log n) steps. You may use one additional
		array of size n as a work space
	*/

	var arr []int = make([]int, 0)
	for i := 1; i <= 5; i++ {
		arr = append(arr, i)
	}
	t := CreateTree(arr)
	fmt.Printf("Tree: %v\n", t.tree)
	t.Update(4, 1000, 0, 0, len(arr)-1)
	fmt.Printf("Tree: %v\n", t.tree)
	fmt.Printf("Query 0: %d\n", t.Query(0, 0, 0, t.elements-1, 0))
	t.Update(3, 9191, 0, 0, len(arr)-1)
	fmt.Printf("Tree: %v\n", t.tree)
	fmt.Printf("Query 3: %d\n", t.Query(0, 3, 0, t.elements-1, 0))
	t.Update(0, 171, 0, 0, len(arr)-1)
	t.Update(5, 880, 0, 0, len(arr)-1)
	fmt.Printf("Tree: %v\n", t.tree)
}
