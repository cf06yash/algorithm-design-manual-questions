package main

import "fmt"

func isIncreasing(arr []int, i int) bool {
	if i == len(arr)-1 {
		return arr[i] > arr[i-1]
	}
	return arr[i] < arr[i+1]
}

func isDecreasing(arr []int, i int) bool {
	if i == len(arr)-1 {
		return arr[i] < arr[i-1]
	}
	return arr[i] > arr[i+1]
}

func isMax(arr []int, i int) bool {
	if i == len(arr)-1 {
		return arr[i] > arr[i-1]
	}
	if i == 0 {
		return arr[i] > arr[i+1]
	}

	return arr[i] > arr[i+1] && arr[i] > arr[i-1]
}

func getMax(arr []int, l, r int) int {
	if len(arr) == 1 {
		return 0
	}
	if l <= r {
		mid := (l + r) / 2
		if isMax(arr, mid) {
			return mid
		} else if isIncreasing(arr, mid) {
			return getMax(arr, mid+1, r)
		} else if isDecreasing(arr, mid) {
			return getMax(arr, l, mid-1)
		}
	}
	return -1
}

func main() {
	/*
		You are given a unimodal array of n distinct elements, meaning that its
		entries are in increasing order up until its maximum element, after which its
		elements are in decreasing order. Give an algorithm to compute the maximum
		element of a unimodal array that runs in O(log n) time.
	*/
	arr := []int{1, 2, 3, 5, 7, 8, 12, 6, 4, 2}
	m := getMax(arr, 0, len(arr)-1)
	if m != -1 {
		fmt.Printf("max elem: %d\n", arr[m])
	} else {
		fmt.Println("error could not find max elem")
	}

}
