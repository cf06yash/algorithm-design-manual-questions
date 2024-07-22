package main

import (
	"fmt"
	"math/rand"
)

func partition(arr *[]int, l, h int) int {
	p := rand.Intn(h-l+1) + l
	(*arr)[p], (*arr)[h] = (*arr)[h], (*arr)[p]
	p = h
	firsthigh := l

	for i := l; i <= h; i++ {
		if (*arr)[i] < (*arr)[p] {
			(*arr)[i], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[i]
			firsthigh++
		}
	}
	(*arr)[p], (*arr)[firsthigh] = (*arr)[firsthigh], (*arr)[p]
	return firsthigh
}

func QuickSelect(arr *[]int, k, l, h int) int {
	if l <= h {
		p := partition(arr, l, h)
		if p == k {
			return (*arr)[p]
		} else if p > k {
			return QuickSelect(arr, k, l, p-1)
		} else {
			return QuickSelect(arr, k, p+1, h)
		}
	}
	return -1
}

func main() {
	/*
		Use the partitioning idea of quicksort to give an algorithm that finds the
		median element of an array of n integers in expected O(n) time. (Hint: must
		you look at both sides of the partition?)
	*/
	var arr []int = []int{1, 3, 2, 2, 3, 1}
	ans := QuickSelect(&arr, len(arr)/2, 0, len(arr)-1)
	fmt.Printf("median : %d\n", ans)
}
