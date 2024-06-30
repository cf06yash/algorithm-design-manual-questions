package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SegmentTree struct {
	tree []int
	n    int
}

type Query struct {
	l int
	r int
}

func minRangeQuery(tree *[]int, low, high, qlow, qhigh, pos int) int {
	if low > high {
		return -1
	}
	//total
	if qlow <= low && qhigh >= high {
		return (*tree)[pos]
	}
	//no overlap
	if qlow > high || qhigh < low {
		return math.MaxInt
	}

	//partial
	mid := (low + high) / 2
	return min(minRangeQuery(tree, low, mid, qlow, qhigh, 2*pos+1), minRangeQuery(tree, mid+1, high, qlow, qhigh, 2*pos+2))
}

func (t *SegmentTree) MinInRange(query *Query) int {
	if query.l < 0 || query.l >= t.n || query.r < 0 || query.r >= t.n {
		return math.MaxInt
	}
	return minRangeQuery(&t.tree, 0, t.n, query.l, query.r, 0)
}

func build(tree *[]int, pos, left, right int, arr *[]int) {
	if left == right {
		(*tree)[pos] = (*arr)[left]
		return
	}
	mid := (left + right) / 2
	build(tree, 2*pos+1, left, mid, arr)
	build(tree, 2*pos+2, mid+1, right, arr)

	(*tree)[pos] = min((*tree)[2*pos+1], (*tree)[2*pos+2])
}

func constructTree(arr *[]int) *SegmentTree {
	var segmentTree *SegmentTree = &SegmentTree{
		tree: make([]int, 4*len(*arr)),
		n:    len(*arr),
	}
	for i := 0; i < len(segmentTree.tree); i++ {
		segmentTree.tree[i] = math.MaxInt
	}
	build(&segmentTree.tree, 0, 0, len(*arr)-1, arr)
	return segmentTree
}

func main() {
	/*
		Suppose that we are given a sequence of n values x1, x2, ..., xn and seek to
		quickly answer repeated queries of the form: given i and j, find the smallest
		value in xi,...,xj .

		(b) Design a data structure that uses O(n) space and answers queries in
		O(log n) time. For partial credit, your data structure can use O(n log n)
		space and have O(log n) query time.
	*/
	var arr []int = []int{1, 3, 2, 7, 9, 11, 4, 5, 6, 8, 10}
	t := constructTree(&arr)
	fmt.Printf("Tree: %v\n", t.tree)
	queries := make([]*Query, 0)

	fmt.Printf("Please enter total number of queries: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	n, _ := strconv.Atoi(temp)
	fmt.Printf("Please enter Queries: %d\n", n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		temp := scanner.Text()
		queryArr := strings.Split(temp, " ")
		l, _ := strconv.Atoi(queryArr[0])
		r, _ := strconv.Atoi(queryArr[1])
		queries = append(queries, &Query{
			l: l,
			r: r,
		})
	}

	fmt.Printf("Running queries...\n")
	for i := 0; i < len(queries); i++ {
		fmt.Printf("Answer for query %v is: %d\n", *queries[i], t.MinInRange(queries[i]))
	}
}
