package main

import "fmt"

func getHIndex(citations []int) int {
	n := len(citations)
	h_index := 0
	counts := make([]int, n+1)

	for i := 0; i < len(citations); i++ {
		if citations[i] > n {
			counts[n]++
		} else {
			counts[i]++
		}
	}
	total := 0
	for i := n; i >= 0; i-- {
		total += counts[i]
		if total >= i {
			h_index = i
			break
		}
	}

	return h_index
}

func main() {
	/*
		Give an efficient algorithm to take the array of citation counts (each count
		is a non-negative integer) of a researcher’s papers, and compute the researcher’s
		h-index. By definition, a scientist has index h if h of his or her n papers have
		been cited at least h times, while the other n−h papers each have no more than
		h citations.
	*/
	var citationsCount []int = []int{51, 6, 721, 1, 23567, 89, 2, 1, 124, 124}
	fmt.Printf("h-index %d\n", getHIndex(citationsCount))

}
