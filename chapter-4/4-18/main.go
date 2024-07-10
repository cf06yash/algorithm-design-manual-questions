package main

import "fmt"

type Heap struct {
	elements []*Trio
}

type Trio struct {
	e, row, col int
}

func (p *Trio) String() string {
	return fmt.Sprintf("%d ", p.e)
}

func (h *Heap) Size() int {
	return len(h.elements)
}

func getLeftChild(idx int) int {
	return 2 * idx
}

func getRightChild(idx int) int {
	return 2*idx + 1
}
func getParent(idx int) int {
	return idx / 2
}

func (h *Heap) bubble_down(idx int) {
	if idx > h.Size()-1 {
		return
	}
	smallest := idx
	left := getLeftChild(idx)
	if left < h.Size() && h.elements[left].e < h.elements[smallest].e {
		smallest = left
	}
	right := getRightChild(idx)
	if right < h.Size() && h.elements[right].e < h.elements[smallest].e {
		smallest = right
	}
	if idx != smallest {
		h.elements[smallest], h.elements[idx] = h.elements[idx], h.elements[smallest]
		h.bubble_down(smallest)
	}
}

func (h *Heap) bubble_up(idx int) {
	if idx > 1 {
		p := getParent(idx)
		if p == 0 {
			return
		} else if h.elements[p].e > h.elements[idx].e {
			h.elements[p], h.elements[idx] = h.elements[idx], h.elements[p]
			h.bubble_up(p)
		}
	}
}

func (h *Heap) Add(e, arrIdx, col int) {
	h.elements = append(h.elements, &Trio{e, arrIdx, col})
	h.bubble_up(h.Size() - 1)
}

func (h *Heap) GetMin() (*Trio, bool) {
	if h.Size() < 2 {
		return nil, false
	}
	m := h.elements[1]
	last := h.elements[h.Size()-1]
	h.elements = h.elements[:h.Size()-1]
	if h.Size() > 1 {
		h.elements[1] = last
		h.bubble_down(1)
	}
	return m, true
}

func GetHeap(arr []*Trio) *Heap {
	h := &Heap{
		elements: make([]*Trio, len(arr)+1),
	}
	for i := 0; i < len(arr); i++ {
		h.elements[i+1] = arr[i]
	}

	firstInternalNode := len(arr) / 2
	for i := firstInternalNode; i > 0; i-- {
		h.bubble_down(i)
	}
	return h
}

func Merge(sortedArrays [][]int, k, n int) []int {
	if len(sortedArrays) == 0 {
		return make([]int, 0)
	}

	res := make([]int, n)
	temp := []*Trio{}
	for i := 0; i < k; i++ {
		temp = append(temp, &Trio{sortedArrays[i][0], i, 0})
	}

	idx := 0
	h := GetHeap(temp)
	filled := 0
	for trio, ok := h.GetMin(); ok; trio, ok = h.GetMin() {
		res[idx] = trio.e
		idx++
		col := trio.col + 1
		row := trio.row
		if col < len(sortedArrays[trio.row]) {
			h.Add(sortedArrays[trio.row][col], row, col)
			filled++
		}
	}

	return res
}

func main() {
	/*
		Give an O(n log k)-time algorithm that merges k sorted lists with a total
		of n elements into one sorted list. (Hint: use a heap to speed up the obvious
		O(kn)-time algorithm).
	*/
	x := [][]int{{1, 5, 9}, {1, 7, 9, 16}, {2, 4, 7, 12, 17}}
	n := 12
	k := 3
	res := Merge(x, k, n)
	fmt.Printf("%v\n", res)
}
