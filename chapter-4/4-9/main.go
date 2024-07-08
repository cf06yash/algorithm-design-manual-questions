package main

import "fmt"

func QuickSort(arr *[]int, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, h)
	}
}

func partition(arr *[]int, l, h int) int {
	p := h
	firsthigh := l

	for i := l; i < h; i++ {
		if (*arr)[i] < (*arr)[p] {
			(*arr)[firsthigh], (*arr)[i] = (*arr)[i], (*arr)[firsthigh]
			firsthigh++
		}
	}
	(*arr)[firsthigh], (*arr)[p] = (*arr)[p], (*arr)[firsthigh]
	return firsthigh
}

func Search(arr []int, key, l, h int) int {
	if l <= h {
		mid := (l + h) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			return Search(arr, key, l, mid-1)
		} else {
			return Search(arr, key, mid+1, h)
		}
	}
	return -1
}

func FindKIntThatSumT(arr []int, res *[]int, l, k, t int) bool {
	if l > len(arr)-1 {
		return false
	}
	if k == 2 {
		for i := l; i < len(arr); i++ {
			x := arr[i]
			y := Search(arr, t-x, i+1, len(arr)-1)
			if y != -1 {
				*res = append(*res, x)
				*res = append(*res, arr[y])
				return true
			}
		}
		return false
	}

	for i := l; i < len(arr); i++ {
		ok := FindKIntThatSumT(arr, res, i+1, k-1, t-arr[i])
		if ok {
			*res = append(*res, arr[i])
			return true
		}
	}

	return false
}

func Sum(res []int) int {
	t := 0
	for i := 0; i < len(res); i++ {
		t += res[i]
	}
	return t
}

func main() {
	arr := []int{5, 4, 2, 1, 8, 7, 6, 10, 3}
	k := 5
	t := 10 + 2 + 7 + 5 + 8
	QuickSort(&arr, 0, len(arr)-1)
	fmt.Printf("Sorted array %d\n", arr)
	res := []int{}
	FindKIntThatSumT(arr, &res, 0, k, t)
	fmt.Printf("nums are %v which equals %d\n", res, Sum(res))
}
