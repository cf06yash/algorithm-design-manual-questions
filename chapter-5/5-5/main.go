package main

import "fmt"

func findItemIdx(arr []int, l int, r int) int {
	if len(arr) == 0 {
		return -1
	}
	if l <= r {
		mid := (l + r) / 2
		if arr[mid] == mid+1 {
			return mid
		} else if arr[mid] < mid+1 {
			return findItemIdx(arr, mid+1, r)
		} else {
			return findItemIdx(arr, l, mid-1)
		}
	}
	return -1
}

func main() {
	/*
		Suppose that you are given a sorted sequence of distinct integers [a1, a2,...,an].
		Give an O(lg n) algorithm to determine whether there exists an index i such that
		ai = i. For example, in [−10, −3, 3, 5, 7], a3 = 3. In [2, 3, 4, 5, 6, 7], there is no
		such i.
	*/

	arr := []int{-10, -3, 3, 5, 7}
	idx := findItemIdx(arr, 0, len(arr)-1)
	if idx != -1 {
		fmt.Printf("Found : item val %d with idx %d\n", arr[idx], idx)
	} else {
		fmt.Printf("Not found\n")
	}
}
