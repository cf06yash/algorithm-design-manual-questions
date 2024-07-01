package main

import (
	"fmt"
)

type SegmentTree struct {
	tree []int
}

func getLeftChild(pos int) int {
	return 2*pos + 1
}

func getRightChild(pos int) int {
	return 2*pos + 2
}

func (t *SegmentTree) build(arr []bool, pos, l, r int) int {
	if l == r {
		if arr[l] {
			t.tree[pos] = 1
		} else {
			t.tree[pos] = 0
		}
		return t.tree[pos]
	}
	mid := (l + r) / 2
	t.tree[pos] = t.build(arr, getLeftChild(pos), l, mid) + t.build(arr, getRightChild(pos), mid+1, r)
	return t.tree[pos]
}

func (t *SegmentTree) queryHelper(l, r, pos, ql, qr int) int {
	if ql <= l && r <= qr {
		return t.tree[pos]
	}
	if ql > r || qr < l {
		return 0
	}
	mid := (l + r) / 2
	return t.queryHelper(l, mid, getLeftChild(pos), ql, qr) + t.queryHelper(mid+1, r, getRightChild(pos), ql, qr)
}

func (t *SegmentTree) Count(n, ql, qr int) int {
	if ql > n-1 || ql < 0 || qr < 0 || qr > n-1 || ql > qr {
		return 0
	}
	return t.queryHelper(0, n-1, 0, ql, qr)
}

func (t *SegmentTree) update(arr *[]bool, idx, l, r, pos int, state bool) {
	if l == r {
		if state {
			t.tree[pos] = 1
		} else {
			t.tree[pos] = 0
		}
		(*arr)[idx] = state
		return
	}
	mid := (l + r) / 2
	if mid < idx {
		t.update(arr, idx, mid+1, r, getRightChild(pos), state)
	} else {
		t.update(arr, idx, l, mid, getLeftChild(pos), state)
	}

	t.tree[pos] = t.tree[getRightChild(pos)] + t.tree[getLeftChild(pos)]
}

func (t *SegmentTree) searchUnoccupied(arr *[]bool, ql, qr, l, r, pos int) int {
	//fmt.Printf("ql: %d,qr: %d,l: %d,r: %d,pos: %d, value: %d\n", ql, qr, l, r, pos, t.tree[pos])
	if ql > r || qr < l {
		return -1
	}
	if t.tree[pos] <= 0 {
		return -1
	}
	if l == r {
		return l
	}
	mid := (l + r) / 2
	left := t.searchUnoccupied(arr, ql, qr, l, mid, getLeftChild(pos))
	if left != -1 {
		return left
	}
	right := t.searchUnoccupied(arr, ql, qr, mid+1, r, getRightChild(pos))
	return right
}

func (t *SegmentTree) Checkin(arr *[]bool, ql int, qr int) int {
	n := len(*arr)
	if ql > n-1 || ql < 0 || qr < 0 || qr > n-1 || ql > qr {
		return -1
	}
	room := t.searchUnoccupied(arr, ql, qr, 0, len(*arr)-1, 0)
	if room > 0 && room < n {
		t.update(arr, room, 0, n-1, 0, false)
	}
	return room
}

func (t *SegmentTree) CheckOut(arr *[]bool, idx int) {
	t.update(arr, idx, 0, len(*arr)-1, 0, true)
}

func CreateNewTree(arr []bool) *SegmentTree {
	tree := &SegmentTree{
		tree: make([]int, 4*len(arr)),
	}

	tree.build(arr, 0, 0, len(arr)-1)
	return tree
}

func main() {
	/*
		You are consulting for a hotel that has n one-bed rooms. When a guest
		checks in, they ask for a room whose number is in the range [l, h]. Propose a
		data structure that supports the following data operations in the allotted time:
		(a) Initialize(n): Initialize the data structure for empty rooms numbered
		1, 2,...,n, in polynomial time.
		(b) Count(l, h): Return the number of available rooms in [l, h], in O(log n)
		time.
		(c) Checkin(l, h): In O(log n) time, return the first empty room in [l, h] and
		mark it occupied, or return NIL if all the rooms in [l, h] are occupied.
		(d) Checkout(x): Mark room x as unoccupied, in O(log n) time.
	*/
	n := 5
	var rooms []bool = make([]bool, n)
	for i := 0; i < n; i++ {
		rooms[i] = true
	}
	t := CreateNewTree(rooms)
	fmt.Printf("tree : %d\n", t.tree)
	fmt.Printf("count 3,3 : %d\n", t.Count(n, 3, 3))
	fmt.Printf("Checking in (1,4) %d\n", t.Checkin(&rooms, 1, 4))
	fmt.Printf("Checking in (1,4) %d\n", t.Checkin(&rooms, 1, 4))
	fmt.Printf("Checking in (1,4) %d\n", t.Checkin(&rooms, 1, 4))
	t.CheckOut(&rooms, 1)
	fmt.Printf("Checking in (1,4) %d\n", t.Checkin(&rooms, 1, 4))
	fmt.Printf("tree : %d\n", t.tree)
	fmt.Printf("count 1,1 : %d\n", t.Count(n, 1, 1))
}
