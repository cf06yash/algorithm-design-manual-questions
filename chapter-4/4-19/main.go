package main

import "fmt"

type Heap struct {
	elements []int
}

func getRightChild(i int) int {
	return 2*i + 1
}

func getLeftChild(i int) int {
	return 2 * i
}

func getParent(i int) int {
	return i / 2
}

func (h *Heap) bubble_down(idx int) {
	if idx > len(h.elements)-1 {
		return
	}
	largest := idx
	left := getLeftChild(idx)
	if left < len(h.elements) && h.elements[left] > h.elements[largest] {
		largest = left
	}
	right := getRightChild(idx)
	if right < len(h.elements) && h.elements[right] > h.elements[largest] {
		largest = right
	}

	if idx != largest {
		h.elements[largest], h.elements[idx] = h.elements[idx], h.elements[largest]
		h.bubble_down(largest)
	}
}

func (h *Heap) GetMax() (int, bool) {
	if len(h.elements) <= 1 {
		return -1, false
	}
	mx := h.elements[1]
	last := h.elements[len(h.elements)-1]
	h.elements[1] = last
	h.elements = h.elements[:len(h.elements)-1]
	h.bubble_down(1)
	return mx, true
}

func (h *Heap) PeekSecondMax() (int, bool) {
	if len(h.elements) <= 2 {
		return -1, false
	}
	rootIdx := 1
	left := getLeftChild(rootIdx)
	right := getRightChild(rootIdx)

	var secondLargest int
	if left < len(h.elements) {
		secondLargest = h.elements[left]
	}

	if right < len(h.elements) && h.elements[right] > secondLargest {
		secondLargest = h.elements[right]
	}

	return secondLargest, true
}

func ConstructHeap(arr []int) (*Heap, bool) {
	if arr == nil || len(arr) == 0 {
		return nil, false
	}
	h := &Heap{
		elements: make([]int, len(arr)+1),
	}
	for i := 0; i < len(arr); i++ {
		h.elements[i+1] = arr[i]
	}

	for i := len(arr) / 2; i >= 1; i-- {
		h.bubble_down(i)
	}
	return h, true
}

func main() {
	/*
		Give an efficient algorithm to find the second-largest key among n keys.
		You can do better than 2n − 3 comparisons

		Then, give an efficient algorithm to find the third-largest key among n keys.
		How many key comparisons does your algorithm do in the worst case? Must your
		algorithm determine which key is largest and second-largest in the process?
	*/
	fmt.Println("go go go")
	var arr = []int{1, 5, 12, 6, 8, 14, 99, 100, 88, 11, 22, 100}
	h, ok := ConstructHeap(arr)
	if !ok {
		fmt.Printf("Error in heap construction!\n")
		return
	}

	secondLargest, ok := h.PeekSecondMax()
	if !ok {
		fmt.Printf("Second largest element not found!\n")
		return
	}
	fmt.Printf("Second largest element is %d\n", secondLargest)
	h.GetMax()
	thirdLargest, ok := h.PeekSecondMax()
	if !ok {
		fmt.Printf("Third largest element not found!\n")
		return
	}
	fmt.Printf("Third largest element is %d\n", thirdLargest)
}
