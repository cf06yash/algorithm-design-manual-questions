package main

import "fmt"

func partition(arr *[]int, l, h int) int {
	firsthigh := h
	p := l

	for i := p + 1; i <= h; i++ {
		if (*arr)[i] > (*arr)[p] {
			(*arr)[i], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[i]
			firsthigh--
		}
	}

	(*arr)[p], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[p]
	return firsthigh
}

func QuickSort(arr *[]int, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, h)
	}
}

func MaxActiveCount(start, end []int) int {
	count := 0
	mc := -1
	i, j := 0, 0
	for i < len(start) && j < len(end) {
		//fmt.Printf("start %d, end %d, count %d\n", start[i], end[j], count)
		if start[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		mc = max(count, mc)
	}

	for i < len(start) {
		count++
		i++
		mc = max(count, mc)
	}

	for j < len(end) {
		count--
		j++
		mc = max(count, mc)
	}

	return mc
}

func main() {

	/*
		A camera at the door tracks the entry time ai and exit time bi (assume
		bi > ai) for each of n persons pi attending a party. Give an O(n log n) algorithm that analyzes this data to determine the time when the most people were
		simultaneously present at the party. You may assume that all entry and exit
		times are distinct (no ties).
	*/

	startTimes := []int{1, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	endTimes := []int{3, 7, 9, 11, 13, 15, 17, 19, 21, 23}

	QuickSort(&startTimes, 0, len(startTimes)-1)
	QuickSort(&endTimes, 0, len(endTimes)-1)

	fmt.Printf("Sorted Array: %v\n", startTimes)
	fmt.Printf("Sorted Array: %v\n", endTimes)

	fmt.Printf("Max Active count in party: %d\n", MaxActiveCount(startTimes, endTimes))
}
