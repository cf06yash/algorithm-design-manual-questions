package main

import (
	"fmt"
	"strings"
)

func checkAnagrams(x, y string) bool {
	fmt.Printf("Checking if strings %s & %s are anagrams...\n", x, y)
	if len(x) != len(y) {
		return false
	}
	x = strings.ToLower(x)
	y = strings.ToLower(y)
	var bucket []int = make([]int, 26)
	for k := 0; k < len(x); k++ {
		i := int(x[k]) - int(rune('a'))
		j := int(y[k]) - int(rune('a'))
		bucket[i]++
		bucket[j]--
	}

	for _, item := range bucket {
		if item != 0 {
			return false
		}
	}
	return true
}

func main() {
	/*
		Two strings X and Y are anagrams if the letters of X can be rearranged
		to form Y . For example, silent/listen, and incest/insect are anagrams. Give an
		efficient algorithm to determine whether strings X and Y are anagrams.
	*/

	var x string = "sileNt"
	var y string = "Listen"

	ok := checkAnagrams(x, y)
	if ok {
		fmt.Println("These strings are anagrams")
	} else {
		fmt.Println("These strings are not anagrams")
	}
}
