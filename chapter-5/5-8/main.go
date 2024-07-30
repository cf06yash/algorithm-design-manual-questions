package main

import "fmt"

func find(arr []int, l, r, key int) int {
	if l >= r {
		return l
	}
	mid := (l + r) / 2
	if arr[mid] > key {
		return find(arr, l, mid-1, key)
	} else {
		return find(arr, mid+1, r, key)
	}
}

func helper(a, b []int, med, l, r int) int {
	if l <= r {
		mid := (l + r) / 2
		posInA := find(a, 0, len(a)-1, b[mid])
		if posInA+mid == med {
			return mid
		} else if posInA+mid < med {
			return helper(a, b, med, mid+1, r)
		} else {
			return helper(a, b, med, l, mid-1)
		}
	}
	return -1
}

func findMedian(a, b []int, med int) int {
	m := len(b)
	n := len(a)
	x := helper(a, b, med, 0, m-1)

	if x == -1 {
		x = helper(b, a, med, 0, n-1)
		return a[x]
	} else {
		return b[x]
	}
}

func main() {
	//Given two sorted arrays A and B of size n and m respectively, find the
	//median of the n + m elements. The overall run time complexity should be
	//O(log(m + n))

	a := []int{1, 2, 3, 4, 5, 6, 7, 23, 56, 99, 100}
	b := []int{2, 3, 4, 5, 8, 11, 17}
	n := len(a)
	m := len(b)
	fmt.Printf("a %v b %v\n", a, b)
	fmt.Printf("a %d b %d\n", n, m)
	if (m+n)%2 == 0 {
		med1 := (m + n) / 2
		med2 := med1 + 1
		median1 := findMedian(a, b, med1)
		fmt.Printf("median 1 element %d\n", median1)

		median2 := findMedian(a, b, med2)
		fmt.Printf("median 2 element %d\n", median2)

		fmt.Printf("final median %f\n", (float64(median1)+float64(median2))/2)
	} else {
		med := (m + n) / 2
		median := findMedian(a, b, med)
		fmt.Printf("median element %d\n", median)
	}

}
