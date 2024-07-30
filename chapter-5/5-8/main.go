package main

import "fmt"

func findMedianOneArray(arr []int, l int) float64 {
	x := getPositionsToFind(0, l)
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += float64(arr[x[i]])
	}

	return sum / float64(len(x))
}

func findPositionsOfElement(arr []int, response *[]int, l, r, key int) {
	if l >= r {
		if arr[l] < key {
			*response = append(*response, l+1)
		} else if arr[l] == key {
			*response = append(*response, l)
			*response = append(*response, l+1)
		} else {
			*response = append(*response, l)
		}
	} else {
		mid := (l + r) / 2
		if arr[mid] == key {
			*response = append(*response, mid)
			findPositionsOfElement(arr, response, l, mid-1, key)
			findPositionsOfElement(arr, response, mid+1, r, key)
		} else if arr[mid] > key {
			findPositionsOfElement(arr, response, l, mid-1, key)
		} else {
			findPositionsOfElement(arr, response, mid+1, r, key)
		}
	}
}

func getMedian(arr1 []int, arr2 []int, pos int) int {
	l := 0
	r := len(arr1) - 1

	for l <= r {
		response := []int{}
		mid := (l + r) / 2
		findPositionsOfElement(arr2, &response, 0, len(arr2)-1, arr1[mid])
		fmt.Printf("response for element %d is %v\n", arr1[mid], response)
		for i := 0; i < len(response); i++ {
			if response[i]+mid == pos {
				return mid
			}
		}

		x := response[len(response)-1]
		if x+mid < pos {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func findPos(arr1 []int, arr2 []int, pos int) float64 {
	idx := getMedian(arr1, arr2, pos)
	if idx != -1 {
		return float64(arr1[idx])
	}
	idx = getMedian(arr2, arr1, pos)
	return float64(arr2[idx])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	if n == 0 && m == 0 {
		return 0
	}
	if n == 0 {
		return findMedianOneArray(nums2, m)
	} else if m == 0 {
		return findMedianOneArray(nums1, n)
	}

	x := getPositionsToFind(n, m)
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += findPos(nums1, nums2, x[i])
	}

	return sum / float64(len(x))
}

func getPositionsToFind(n, m int) []int {
	x := []int{}
	if n+m == 0 {
		return x
	}
	total := n + m
	if total%2 == 0 {
		a := (total / 2) - 1
		x = append(x, a)
		x = append(x, a+1)
	} else {
		x = append(x, total/2)
	}
	return x
}

func main() {
	//Given two sorted arrays A and B of size n and m respectively, find the
	//median of the n + m elements. The overall run time complexity should be
	//O(log(m + n))

	a := []int{1, 2, 3, 4, 5, 6, 7, 23, 56, 99, 100}
	b := []int{2, 3, 4, 5, 8, 11, 17}
	median := findMedianSortedArrays(a, b)
	fmt.Printf("median %f\n", median)

}
