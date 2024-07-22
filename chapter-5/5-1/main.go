package main

import "fmt"

func getLargestHelper(arr []int, l int, r int) int {
	if l <= r {
		mid := (l + r) / 2
		if mid < len(arr)-1 && arr[mid+1] < arr[mid] {
			return mid
		} else if mid == len(arr)-1 {
			return mid
		} else if arr[mid] >= arr[l] {
			return getLargestHelper(arr, mid+1, r)
		} else {
			return getLargestHelper(arr, l, mid-1)
		}
	}
	return -1
}

func getLargestFromRotatedArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	index := getLargestHelper(arr, 0, len(arr)-1)
	fmt.Printf("Found index %d\n", index)
	if index == -1 {
		return -1
	}
	return arr[index]
}

func main() {
	/*
		Suppose you are given a sorted array A of size n that has been circularly
		shifted k positions to the right. For example, [35, 42, 5, 15, 27, 29] is a sorted
		array that has been circularly shifted k = 2 positions, while [27, 29, 35, 42, 5, 15]
		has been shifted k = 4 positions

		Suppose you do not know what k is. Give an O(lg n) algorithm to find the
		largest number in A. For partial credit, you may give an O(n) algorithm
	*/

	arr := []int{3, 3, 3, 1, 2}
	largestNum := getLargestFromRotatedArray(arr)
	fmt.Printf("largest: %d\n", largestNum)
}
