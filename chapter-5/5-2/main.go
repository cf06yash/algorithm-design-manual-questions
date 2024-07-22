package main

import "fmt"

func findMissing(arr []int, l, r int) int {
	if l == r {
		if l+1 == arr[l] {
			return -1
		} else if l+1 < arr[l] {
			return l + 1
		}
	}
	if l < r {
		mid := (l + r) / 2
		if mid+1 == arr[mid] && mid < len(arr)-1 && arr[mid+1]-arr[mid] >= 2 {
			return mid + 2
		} else if mid+1 == arr[mid] {
			return findMissing(arr, mid+1, r)
		} else if mid+1 < arr[mid] {
			return findMissing(arr, l, mid-1)
		}
	}

	return -1
}

func main() {
	/*
		A sorted array of size n contains distinct integers between 1 and n + 1, with
		one element missing. Give an O(log n) algorithm to find the missing integer,
		without using any extra space.
	*/

	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	missing := findMissing(arr, 0, len(arr)-1)
	fmt.Printf("missing: %d\n", missing)
}
