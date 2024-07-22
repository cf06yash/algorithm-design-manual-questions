package main

import "fmt"

func Reverse(arr *[]int, i, j int) {
	if j >= len(*arr) {
		return
	}
	if i >= len(*arr) {
		return
	}

	for ; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func partition(arr *[]int, l, h int) int {
	p := h
	firsthigh := l

	for i := l; i < p; i++ {
		if (*arr)[i] < (*arr)[p] {
			Reverse(arr, firsthigh, i)
			firsthigh++
		}
	}

	Reverse(arr, firsthigh, p)
	return firsthigh

}

func Sort(arr *[]int, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		fmt.Printf("returned partition %d\n", p)
		Sort(arr, l, p-1)
		Sort(arr, p+1, h)
	}
}

func main() {
	/*
		Suppose you are given a permutation p of the integers 1 to n, and seek
		to sort them to be in increasing order [1,...,n]. The only operation at your
		disposal is reverse(p,i,j), which reverses the elements of a subsequence pi,...,pj
		in the permutation. For the permutation [1, 4, 3, 2, 5] one reversal (of the second
		through fourth elements) suffices to sort.
		• Show that it is possible to sort any permutation using O(n) reversals.
		• Now suppose that the cost of reverse(p,i,j) is equal to its length, the number of elements in the range, |j − i| + 1. Design an algorithm that sorts p
		in O(n log2 n) cost. Analyze the running time and cost of your algorithm
		and prove correctness
	*/

	var arr = []int{6, 7, 8, 2, 1, 4, 3, 5, 9, 32, 45, -12, 66}
	Sort(&arr, 0, len(arr)-1)
	fmt.Printf("%v\n", arr)
}
