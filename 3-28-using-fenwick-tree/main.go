package main

import (
	"fmt"
)

type FenwickTree struct {
	tree []int
}

func getNext(i int) int {
	//2s compliment
	x := ^i + 1
	//& with original
	y := i & x
	// add to original number
	return i + y
}

func getParent(i int) int {
	x := ^i + 1
	//& with original
	y := i & x
	// substract to original number
	return i - y
}

func (t *FenwickTree) add(arr []int, idx, treeIdx int) {
	if treeIdx > len(t.tree)-1 {
		return
	}
	t.tree[treeIdx] += arr[idx]
	nextIdx := getNext(treeIdx)
	t.add(arr, idx, nextIdx)
}

func (t *FenwickTree) prefixSumQueryHelper(treeIdx int) int {
	if treeIdx <= 0 {
		return 0
	}
	return t.tree[treeIdx] + t.prefixSumQueryHelper(getParent(treeIdx))
}

func (t *FenwickTree) PrefixSumQuery(rightIdx int) int {
	return t.prefixSumQueryHelper(rightIdx + 1)
}

func (t *FenwickTree) updateHelper(treeIdx int, value int) {
	if treeIdx > len(t.tree)-1 {
		return
	}
	t.tree[treeIdx] += value
	nextIdx := getNext(treeIdx)
	t.updateHelper(nextIdx, value)
}

func (t *FenwickTree) Update(idx int, value int) {
	t.updateHelper(idx+1, value)
}

func CreateTree(arr []int) *FenwickTree {
	n := len(arr)
	t := &FenwickTree{
		tree: make([]int, n+1),
	}
	for i := 0; i < len(arr); i++ {
		t.add(arr, i, i+1)
	}
	return t
}

func main() {
	/*
		arr := make([]int, 0)
		for i := 0; i < 10; i++ {
			arr = append(arr, i*rand.Intn(12))
		}
	*/

	/*
			Let A[1..n] be an array of real numbers. Design an algorithm to perform
		any sequence of the following operations:

		• Add(i,y) – Add the value y to the ith number.
		• Partial-sum(i) – Return the sum of the first i numbers, that is, i
		j=1 A[j].
		There are no insertions or deletions; the only change is to the values of the numbers. Each operation should take O(log n) steps. You may use one additional
		array of size n as a work space
	*/
	arr := []int{3, 2, -1, 6, 5, 4, -3, 3, 7, 2, 3}
	t := CreateTree(arr)
	fmt.Printf("Fenwick tree: %v\n", t.tree)
	fmt.Printf("Query (0,7): %d\n", t.PrefixSumQuery(7))
	t.Update(8, -19)
	t.Update(7, -15)
	t.Update(2, -3)
	fmt.Printf("Query (0,7): %d\n", t.PrefixSumQuery(7))
}
