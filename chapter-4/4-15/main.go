package main

import "fmt"

type Pair struct {
	start, end int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d,%d)", p.start, p.end)
}

func compare(a, b *Pair) int {
	if a.start < b.start {
		return -1
	} else if a.start > b.start {
		return 1
	} else if a.end <= b.end {
		return -1
	}
	return 1
}

func partition(intervals *[]*Pair, l, h int) int {
	p := l
	firsthigh := h

	for i := h; i > l; i-- {
		if compare((*intervals)[i], (*intervals)[p]) > 0 {
			(*intervals)[i], (*intervals)[firsthigh] = (*intervals)[firsthigh], (*intervals)[i]
			firsthigh--
		}
	}
	(*intervals)[p], (*intervals)[firsthigh] = (*intervals)[firsthigh], (*intervals)[p]
	return firsthigh
}

func QuickSort(intervals *[]*Pair, l, h int) {
	if l < h {
		p := partition(intervals, l, h)
		QuickSort(intervals, l, p-1)
		QuickSort(intervals, p+1, h)
	}
}

func main() {
	/*
		You are given a set S of n intervals on a line, with the ith interval described
		by its left and right endpoints (li, ri). Give an O(n log n) algorithm to identify
		a point p on the line that is in the largest number of intervals.
		As an example, for S = {(10, 40),(20, 60),(50, 90),(15, 70)} no point exists in
		all four intervals, but p = 50 is an example of a point in three intervals. You
		can assume an endpoint counts as being in its interval.
	*/

	intervals := []*Pair{{1, 5}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}

	QuickSort(&intervals, 0, len(intervals)-1)
	fmt.Printf("Sorted Intervals: %v\n", intervals)

	if len(intervals) == 0 {
		return
	}
	var maxPoint = intervals[0].start
	var maxCount = 1
	count := 0
	i, j := 0, 0
	for i < len(intervals) && j < len(intervals) {
		var point = 0
		if intervals[i].start <= intervals[j].end {
			point = intervals[i].start
			count++
			i++
		} else if intervals[i].start > intervals[j].end {
			point = intervals[i].end
			count--
			j++
		}
		if maxCount < count {
			maxPoint = point
			maxCount = count
		}
	}

	fmt.Printf("Point with max overlapping intervals: %d and total overlapping intervals %d\n", maxPoint, maxCount)
}
