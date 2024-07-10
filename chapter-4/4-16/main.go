package main

import (
	"fmt"
	"math"
	"sort"
)

type Segment struct {
	Start, End int
}

func (seg *Segment) Contains(key int) bool {
	return key >= seg.Start && key <= seg.End
}

func Compare(a, b *Segment) int {
	if a.Start < b.Start {
		return -1
	} else if a.Start > b.Start {
		return 1
	} else if a.End <= b.End {
		return -1
	} else {
		return 1
	}
}

type SegmentSlice []*Segment

func (s SegmentSlice) Len() int           { return len(s) }
func (s SegmentSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SegmentSlice) Less(i, j int) bool { return Compare(s[i], s[j]) < 0 }

func Search(segments []*Segment, res *[]int, key, l, h int) {
	if l <= h {
		mid := (l + h) / 2
		if segments[mid].Contains(key) {
			*res = append(*res, mid)
			Search(segments, res, key, l, mid-1)
			Search(segments, res, key, mid+1, h)
			return
		} else if segments[mid].Start > key {
			Search(segments, res, key, l, mid-1)
		} else if segments[mid].End < key {
			Search(segments, res, key, mid+1, h)
		}
	}
}

func BuildPathToM(segments []*Segment, start *Segment, startIdx, m int) (int, bool) {
	var count = 1
	for !start.Contains(m) {
		res := []int{}
		Search(segments, &res, start.End, startIdx+1, len(segments)-1)
		if len(res) == 0 {
			return -1, false
		}

		var nextSegment *Segment
		nextIdx := startIdx
		nextEnd := start.End
		for i := 0; i < len(res); i++ {
			next := segments[res[i]]
			if next.End > nextEnd {
				nextSegment = next
				nextIdx = res[i]
			}
		}
		if nextSegment == nil {
			return -1, false
		}
		start = nextSegment
		startIdx = nextIdx
		count++
	}

	return count, true

}

func GetPathsThatStartAtZeroAndReachM(segments []*Segment, m int) map[*Segment]int {
	counter := make(map[*Segment]int)
	if len(segments) == 0 {
		return counter
	}
	sort.Sort(SegmentSlice(segments))
	zeroSegments := []int{}
	Search(segments, &zeroSegments, 0, 0, len(segments)-1)
	for i := 0; i < len(zeroSegments); i++ {
		//counter[zeroSegments[i]] = -1
		start := segments[zeroSegments[i]]
		count, ok := BuildPathToM(segments, start, zeroSegments[i], m)
		if ok {
			counter[start] = count
		}
	}

	return counter
}

func main() {
	/*
		You are given a set S of n segments on the line, where segment Si ranges
		from li to ri. Give an efficient algorithm to select the fewest number of segments
		whose union completely covers the interval from 0 to m.
	*/

	m := 20
	segments := []*Segment{
		{13, 18},
		{16, 20},
		{-11, -1},
		{0, 4},
		{0, 16},
		{-2, 3},
		{2, 13},
		{5, 11},
		{10, 15},
	}
	counter := GetPathsThatStartAtZeroAndReachM(segments, m)

	var minCount = math.MaxInt
	var ans *Segment
	for key, value := range counter {
		fmt.Printf("segment %v reaches %d in %d unions\n", key, m, value)
		if value < minCount {
			ans = key
			minCount = value
		}
	}

	fmt.Printf("Segment that takes smallest path is %v and it reaches in %d steps\n", ans, minCount)
}
