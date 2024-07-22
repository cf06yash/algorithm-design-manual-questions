package main

import (
	"fmt"
)

var testCases = []struct {
	a        []int
	m        int
	expected int
}{
	{
		a:        []int{1, 2, 4, 5, 7},
		m:        8,
		expected: 3,
	},
	{
		a:        []int{2, 3, 4, 5},
		m:        6,
		expected: 1,
	},
	{
		a:        []int{1, 2, 3, 4, 5},
		m:        6,
		expected: 6,
	},
	{
		a:        []int{1, 10, 20, 30, 40},
		m:        50,
		expected: 2,
	},
	{
		a:        []int{2},
		m:        5,
		expected: 1,
	},
	{
		a:        []int{},
		m:        5,
		expected: 1,
	},
	{
		a:        []int{1, 3, 5, 7, 9},
		m:        1000000,
		expected: 2,
	},
	{
		a:        []int{5, 6, 7, 8, 9},
		m:        10,
		expected: 1,
	},
}

func runTestCases() {
	fmt.Println("Running test cases:")
	for i, tc := range testCases {
		result := findSmallestMissing(tc.a, 0, len(tc.a)-1, tc.m)
		fmt.Printf("Test case %d: Expected %d, Got %d\n", i, tc.expected, result)
	}
}

func findSmallestMissing(arr []int, l, r, m int) int {
	if len(arr) == 0 {
		return 1
	}
	if l >= r {
		if arr[l] == l+1 && l+2 <= m {
			return l + 2
		} else if arr[l] == l+1 {
			return -1
		} else if arr[l] > l+1 {
			return l + 1
		}
	}
	if l < r {
		mid := (l + r) / 2
		if arr[mid] == mid+1 {
			return findSmallestMissing(arr, mid+1, r, m)
		} else if arr[mid] > mid+1 {
			return findSmallestMissing(arr, l, mid-1, m)
		}
	}

	return -1
}

func main() {
	/*
		Suppose that you are given a sorted sequence of distinct integers a =
		[a1, a2,...,an], drawn from 1 to m where n<m. Give an O(lg n) algorithm to
		find an integer ≤ m that is not present in a. For full credit, find the smallest
		such integer x such that 1 ≤ x ≤ m.
	*/
	m := 40
	arr := []int{1, 2, 3, 5, 12, 36}
	num := findSmallestMissing(arr, 0, len(arr)-1, m)
	fmt.Printf("Found number %d\n", num)
	runTestCases()
}
