package main

import "fmt"

type Pair struct {
	s, e int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d,%d) ", p.s, p.e)
}

func compare(a, b *Pair) int {
	if a.s < b.s {
		return -1
	} else if a.s > b.s {
		return 1
	} else if a.e <= b.e {
		return -1
	} else {
		return 1
	}
}

func partition(arr *[]*Pair, l, h int) int {
	p := l
	firsthigh := h

	for i := h; i > l; i-- {
		if compare((*arr)[i], (*arr)[p]) > 0 {
			(*arr)[i], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[i]
			firsthigh--
		}
	}
	(*arr)[p], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[p]
	return firsthigh
}

func QuickSort(arr *[]*Pair, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, h)
	}
}

func Overlap(a, b *Pair) bool {
	return !(b.e < a.s || b.s > a.e)
}

func Merge(intervals []*Pair) []*Pair {
	QuickSort(&intervals, 0, len(intervals)-1)
	fmt.Printf("Sorted array: %v\n", intervals)

	cur := intervals[0]
	i := 0
	res := []*Pair{}
	for i < len(intervals)-1 {
		next := intervals[i+1]
		if Overlap(cur, next) {
			cur = &Pair{min(cur.s, next.s), max(cur.e, next.e)}
			i++
		} else {
			res = append(res, cur)
			cur = next
			i++
		}
	}

	if cur != nil {
		res = append(res, cur)
	}
	return res
}

func main() {
	/*
		Given a list I of n intervals, specified as (xi, yi) pairs, return a list where
		the overlapping intervals are merged. For I = {(1, 3),(2, 6),(8, 10),(7, 18)} the
		output should be {(1, 6),(7, 18)}. Your algorithm should run in worst-case
		O(n log n) time complexity.
	*/
	intervals := []*Pair{{7, 7}, {2, 6}, {8, 10}, {16, 23}, {23, 28}, {24, 25}, {1, 3}}
	fmt.Printf("intervals: %v\n", intervals)
	mergedIntervals := Merge(intervals)
	fmt.Printf("merged intervals %v\n", mergedIntervals)
}
