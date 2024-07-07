package main

import "fmt"

func main() {
	/*
		Let S be a sorted array of n integers. Give an algorithm that finds the pair
		x, y ∈ S that maximizes |x − y|. Your algorithm must run in O(1) worst-case time.
	*/
	s := []int{-11, -9, 2, 34, 99, 101}
	fmt.Printf("For sorted array max |x-y| would be for 0th and (n-1)th index. For this array they are: %d & %d\n", s[0], s[len(s)-1])
}
