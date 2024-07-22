package main

import (
	"fmt"
	"math/rand"
)

func partition(arr []int, l, h int) int {
	p := rand.Intn(h-l+1) + l
	arr[p], arr[h] = arr[h], arr[p]
	p = h
	fh := l

	for i := l; i < h; i++ {
		if arr[i] < arr[p] {
			arr[i], arr[fh] = arr[fh], arr[i]
			fh++
		}
	}
	arr[p], arr[fh] = arr[fh], arr[p]
	return fh
}

func QuickSelect(nums []int, l, h, key int) int {
	if l <= h {
		p := partition(nums, l, h)
		if p == key {
			return nums[p]
		} else if p < key {
			return QuickSelect(nums, p+1, h, key)
		} else {
			return QuickSelect(nums, l, p-1, key)
		}
	}
	return -1
}

func getMedian(nums []int) int {
	return QuickSelect(nums, 0, len(nums)-1, len(nums)/2)
}

func wiggleSort(nums []int) {
	median := getMedian(nums)
	fmt.Printf("median %d\n", median)

	n := len(nums)
	i, j, k := 0, 0, n-1

	for j <= k {
		if nums[Indexer(j, n)] > median {
			nums[Indexer(i, n)], nums[Indexer(j, n)] = nums[Indexer(j, n)], nums[Indexer(i, n)]
			i++
			j++
		} else if nums[Indexer(j, n)] < median {
			nums[Indexer(k, n)], nums[Indexer(j, n)] = nums[Indexer(j, n)], nums[Indexer(k, n)]
			k--
		} else {
			j++
		}
	}
}

func Indexer(i, n int) int {
	return (1 + 2*i) % (n | 1)
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Printf("wiggle sort: %v\n", nums)

	nums = []int{1, 2, 2, 3}
	wiggleSort(nums)
	fmt.Printf("wiggle sort: %v\n", nums)
}
