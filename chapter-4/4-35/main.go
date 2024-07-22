package main

import (
	"fmt"
	"math"
	"sort"
)

type Queue struct {
	elems []int
}

func (q *Queue) IsEmpty() bool {
	return q.elems == nil || len(q.elems) == 0
}

func (q *Queue) Enqueue(a int) {
	q.elems = append(q.elems, a)
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.elems[0], true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	ans := q.elems[0]
	q.elems = q.elems[1:]
	return ans, true
}

func merge(arr []int, al, ah, bl, bh int) {
	if len(arr) == 0 {
		return
	}
	q1 := &Queue{
		elems: make([]int, 0),
	}
	q2 := &Queue{
		elems: make([]int, 0),
	}

	for i := al; i <= ah; i++ {
		q1.Enqueue(arr[i])
	}
	for i := bl; i <= bh; i++ {
		q2.Enqueue(arr[i])
	}
	k := 0

	for !q1.IsEmpty() && !q2.IsEmpty() {
		e1, _ := q1.Peek()
		e2, _ := q2.Peek()
		if e1 <= e2 {
			arr[k] = e1
			k++
			q1.Dequeue()
		} else {
			arr[k] = e2
			k++
			q2.Dequeue()
		}
	}

	for !q1.IsEmpty() {
		arr[k], _ = q1.Dequeue()
		k++
	}

	for !q2.IsEmpty() {
		arr[k], _ = q2.Dequeue()
		k++
	}
}

func main() {
	/*
		Let A[1..n] be an array such that the first n−√n elements are already sorted
		(though we know nothing about the remaining elements). Give an algorithm
		that sorts A in substantially better than n log n steps.
	*/
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 16, 11, 14, 13, 12}
	n := len(arr)
	unsorted := int(math.Sqrt(float64(n)))

	i := n - unsorted
	j := n - 1
	fmt.Printf("n %d idx start %d idx end %d\n", n, i, j)

	sort.Slice(arr[i:j+1], func(x, y int) bool {
		return arr[i+x] < arr[i+y]
	}) // this can be done linear by counting sort

	merge(arr, 0, i-1, i, j)
	fmt.Printf("sorted arr %v\n", arr)
}
