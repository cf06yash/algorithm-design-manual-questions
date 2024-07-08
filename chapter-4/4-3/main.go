package main

import "fmt"

type Pair struct {
	a, b int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d,%d) ", p.a, p.b)
}

func QuickSort(arr *[]int, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, h)
	}
}

func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

func partition(arr *[]int, l, h int) int {
	firsthigh := l
	p := h

	for i := 0; i < h; i++ {
		if (*arr)[i] < (*arr)[p] {
			swap(arr, i, firsthigh)
			firsthigh++
		}
	}

	return firsthigh
}

func GetPairs(arr []int) []*Pair {
	ans := []*Pair{}
	QuickSort(&arr, 0, len(arr)-1)
	for i := 0; i < len(arr)/2; i++ {
		ans = append(ans, &Pair{
			a: arr[i],
			b: arr[len(arr)-i-1],
		})
	}

	return ans
}

func main() {
	/*
		Take a list of 2n real numbers as input. Design an O(n log n) algorithm
		that partitions the numbers into n pairs, with the property that the partition
		minimizes the maximum sum of a pair. For example, say we are given the
		numbers (1,3,5,9). The possible partitions are ((1,3),(5,9)), ((1,5),(3,9)), and
		((1,9),(3,5)). The pair sums for these partitions are (4,14), (6,12), and (10,8).
		Thus, the third partition has 10 as its maximum sum, which is the minimum
		over the three partitions.
	*/
	arr := []int{3, 21, 11, 16, 96, 96}
	ans := GetPairs(arr)

	fmt.Printf("Generated Pairs are: %v\n", ans)
}
