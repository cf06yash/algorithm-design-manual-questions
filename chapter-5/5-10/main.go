package main

import "fmt"

func getSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func getCountOfPeices(arr []int, length int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count += arr[i] / length
	}
	return count
}

func getMaxPieceLength(arr []int, k int) int {
	total := getSum(arr)
	l := 1
	r := total / k
	for l <= r {
		mid := (l + r) / 2
		c := getCountOfPeices(arr, mid)
		if k <= c {
			l = mid + 1
		} else if k > c {
			r = mid - 1
		}
	}

	return r
}

func main() {
	//We are given n wooden sticks, each of integer length, where the ith piece
	//has length L[i]. We seek to cut them so that we end up with k pieces of exactly
	//the same length, in addition to other fragments. Furthermore, we want these k
	//pieces to be as large as possible
	//Give a correct and efficient algorithm that, for a given L and k, returns the
	//maximum possible length of the k equal pieces cut from the initial n sticks.

	arr := []int{10, 6, 5, 3}
	k := 4
	length := getMaxPieceLength(arr, k)
	fmt.Printf("max length is %d for %d pieces for arr %v\n", length, k, arr)
}
