package main

import (
	"fmt"
	"math"
)

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(a int) {
	if q.elements == nil {
		q.elements = make([]int, 0)
	}
	q.elements = append(q.elements, a)
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.elements[0], true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	elem := q.elements[0]
	q.elements = q.elements[1:]
	return elem, true
}

func (q *Queue) IsEmpty() bool {
	return q.elements == nil || len(q.elements) == 0
}

func CreateNewQueue() *Queue {
	return &Queue{
		elements: make([]int, 0),
	}
}

func mergeSort(A *[]int, l, h int) {
	if l < h {
		mid := (l + h) / 2
		mergeSort(A, l, mid)
		mergeSort(A, mid+1, h)
		merge(A, l, mid, h)
	}
}

func merge(A *[]int, l, mid, h int) {
	q1 := CreateNewQueue()
	q2 := CreateNewQueue()

	for i := l; i <= mid; i++ {
		q1.Enqueue((*A)[i])
	}
	for i := mid + 1; i <= h; i++ {
		q2.Enqueue((*A)[i])
	}

	i := l
	for !q1.IsEmpty() && !q2.IsEmpty() {
		e1, _ := q1.Peek()
		e2, _ := q2.Peek()

		if e1 <= e2 {
			(*A)[i] = e1
			q1.Dequeue()
		} else {
			(*A)[i] = e2
			q2.Dequeue()
		}
		i++
	}

	for !q1.IsEmpty() {

		e1, _ := q1.Peek()
		(*A)[i] = e1
		q1.Dequeue()
		i++
	}

	for !q2.IsEmpty() {

		e2, _ := q2.Peek()
		(*A)[i] = e2
		q2.Dequeue()
		i++
	}
}

func main() {
	/*
		Let S be an unsorted array of n integers. Give an algorithm that finds the
		pair x, y ∈ S that minimizes |x − y|, for x != y. Your algorithm must run in
		O(n log n) worst-case time
	*/
	arr := []int{3, 21, 11, 16, 96, 96}
	mergeSort(&arr, 0, len(arr)-1)
	fmt.Printf("sorted array: %v\n", arr)

	minDiff := math.MaxInt
	px := -1
	py := -1
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] != arr[i+1]) && (int(math.Abs(float64(arr[i]-arr[i+1]))) < minDiff) {
			minDiff = int(math.Abs(float64(arr[i] - arr[i+1])))
			px = arr[i]
			py = arr[i+1]
		}
	}

	fmt.Printf("Minimum |x-y| where x!=y is %d where x = %d & y = %d\n", minDiff, px, py)
}
