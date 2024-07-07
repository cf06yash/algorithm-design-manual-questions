package main

import (
	"fmt"
	"math"
)

func Solve(s []int) (int, [2]int, bool) {
	ans := [2]int{}
	if len(s) == 0 {
		return math.MaxInt, ans, false
	}
	min_elem := s[0]
	max_elem := s[0]
	for _, i := range s {
		min_elem = min(min_elem, i)
		max_elem = max(max_elem, i)
		ans[0] = max_elem
		ans[1] = min_elem
	}

	return int(math.Abs(float64(max_elem) - float64(min_elem))), ans, true
}

func main() {
	/*
		Let S be an unsorted array of n integers. Give an algorithm that finds the
		pair x, y ∈ S that maximizes |x−y|. Your algorithm must run in O(n) worst-case
		time.
	*/
	s := []int{6, 13, 19, 3, 8, -31, -1}
	value, ans, ok := Solve(s)
	if ok {
		fmt.Printf("|x−y| is %d for values %d & %d \n", value, ans[0], ans[1])
	} else {
		fmt.Printf("Empty array!")
	}

}
