package main

import (
	"fmt"
	"math"
)

type Pair struct {
	a, b int
}

type Queue struct {
	elements []int
}

func (q *Queue) IsEmpty() bool {
	return q.elements == nil || len(q.elements) == 0
}

func (q *Queue) Enqueue(a int) {
	if q.IsEmpty() {
		q.elements = make([]int, 0)
	}
	q.elements = append(q.elements, a)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	elem := q.elements[0]
	q.elements = q.elements[1:]
	return elem, true
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.elements[0], true
}

func CreateNewQueue() *Queue {
	return &Queue{
		elements: make([]int, 0),
	}
}

func MergeSort(arr *[]int, l, h int) {
	if l < h {
		mid := (l + h) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, h)
		merge(arr, l, mid, h)
	}
}

func merge(arr *[]int, l, mid, h int) {
	q1 := CreateNewQueue()
	q2 := CreateNewQueue()

	for i := l; i <= mid; i++ {
		q1.Enqueue((*arr)[i])
	}
	for i := mid + 1; i <= h; i++ {
		q2.Enqueue((*arr)[i])
	}

	i := l
	for !q1.IsEmpty() && !q2.IsEmpty() {
		e1, _ := q1.Peek()
		e2, _ := q2.Peek()

		if e1 <= e2 {
			(*arr)[i] = e1
			q1.Dequeue()
		} else {
			(*arr)[i] = e2
			q2.Dequeue()
		}
		i++
	}

	for !q1.IsEmpty() {
		e1, _ := q1.Dequeue()
		(*arr)[i] = e1
		i++
	}

	for !q2.IsEmpty() {
		e2, _ := q2.Dequeue()
		(*arr)[i] = e2
		i++
	}
}

func Search(arr []int, l, h, key int) int {
	if l < h {
		mid := (l + h) / 2
		if arr[mid] > key {
			return Search(arr, l, mid-1, key)
		} else if arr[mid] < key {
			return Search(arr, mid+1, h, key)
		}
		return mid
	}
	return -1
}

func findPair(s1, s2 []int, x int) *Pair {
	p := &Pair{
		a: math.MaxInt, b: math.MaxInt,
	}

	for i := 0; i < len(s1); i++ {
		key := x - s1[i]
		j := Search(s2, 0, len(s2)-1, key)
		if j >= 0 {
			p.a = s1[i]
			p.b = s2[j]
			return p
		}
	}

	return p
}

func main() {
	/*
	   Given two sets S1 and S2 (each of size n), and a number x, describe an
	   O(n log n) algorithm for finding whether there exists a pair of elements, one
	   from S1 and one from S2, that add up to x. (For partial credit, give a Θ(n2)
	   algorithm for this problem.)
	*/
	s1 := []int{5, 2, 11, 8, 15}
	s2 := []int{9, 3, 10, 7, 13}
	x := 20
	n := 5

	if len(s1) == 0 || len(s2) == 0 {
		return
	}

	MergeSort(&s2, 0, n-1)

	p := findPair(s1, s2, x)
	if p.a != math.MaxInt && p.b != math.MaxInt {
		fmt.Printf("Found Pair (%d,%d)\n", p.a, p.b)
	} else {
		fmt.Println("No pair found")
	}
}
