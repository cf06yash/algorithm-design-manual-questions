package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Let S be a sorted array of n integers. Give an algorithm that finds the pair
		x, y ∈ S that minimizes |x − y|, for x != y. Your algorithm must run in O(n)
		worst-case time.
	*/
	arr := []int{3, 11, 16, 21, 96, 96}
	minDiff := math.MaxInt
	px := -1
	py := -1
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] != arr[i+1]) && (int(math.Abs(float64(arr[i]-arr[i+1]))) < minDiff) {
			minDiff = int(math.Abs(float64(arr[i] - arr[i+1])))
			px = arr[i]
			py = arr[i+1]
		}
	}

	fmt.Printf("Minimum |x-y| where x!=y is %d where x = %d & y = %d\n", minDiff, px, py)
}
