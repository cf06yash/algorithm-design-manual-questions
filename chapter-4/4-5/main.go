package main

import (
	"fmt"
	"math"
)

func Mode(arr []int) int {
	var counter map[int]int = make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		old, ok := counter[arr[i]]
		if ok {
			counter[arr[i]] = old + 1
		} else {
			counter[arr[i]] = 1
		}
	}

	mc := -1
	mv := math.MaxInt
	for k, v := range counter {
		if v > mc {
			mv = k
		}
	}

	return mv
}

func main() {
	/*
		The mode of a bag of numbers is the number that occurs most frequently in
		the set. The set {4, 6, 2, 4, 3, 1} has a mode of 4. Give an efficient and correct
		algorithm to compute the mode of a bag of n numbers.
	*/
	arr := []int{4, 6, 2, 4, 3, 1}
	fmt.Printf("Mode is %d\n", Mode(arr))
}
