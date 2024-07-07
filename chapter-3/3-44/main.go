package main

import "fmt"

type SegmentTree struct {
	tree  []int
	elems int
}

func getLeftChild(pos int) int {
	return 2*pos + 1
}

func getRightChild(pos int) int {
	return 2*pos + 2
}

func (st *SegmentTree) build(x []int, l, r, pos int) int {
	if l == r {
		st.tree[pos] = x[l]
		return st.tree[pos]
	}
	mid := (l + r) / 2
	st.tree[pos] = st.build(x, l, mid, getLeftChild(pos)) * st.build(x, mid+1, r, getRightChild(pos))
	return st.tree[pos]
}

func (st *SegmentTree) Search(ql, qr, l, r, pos int) int {
	if qr < l || r < ql {
		return 1
	}
	if ql <= l && r <= qr {
		return st.tree[pos]
	}
	mid := (l + r) / 2
	return st.Search(ql, qr, l, mid, getLeftChild(pos)) * st.Search(ql, qr, mid+1, r, getRightChild(pos))
}

func CreateSegmentTree(x []int) *SegmentTree {
	st := &SegmentTree{
		tree:  make([]int, 4*len(x)),
		elems: len(x),
	}
	st.build(x, 0, len(x)-1, 0)
	return st
}

func main() {
	/*
		You have an unordered array X of n integers. Find the array M containing
		n elements where Mi is the product of all integers in X except for Xi. You may
		not use division. You can use extra memory. (Hint: there are solutions faster
		than O(n2).)
	*/
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := make([]int, len(x))
	st := CreateSegmentTree(x)

	for i := 0; i < len(x); i++ {
		if i == 0 {
			m[i] = st.Search(1, len(x)-1, 0, len(x)-1, 0)
		} else if i == len(x)-1 {
			m[i] = st.Search(0, len(x)-2, 0, len(x)-1, 0)
		} else {
			m[i] = st.Search(0, i-1, 0, len(x)-1, 0) * st.Search(i+1, len(x)-1, 0, len(x)-1, 0)
		}
	}

	fmt.Printf("%v\n", m)
}
