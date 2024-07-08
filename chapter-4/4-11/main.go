package main

import "fmt"

func GetElements(arr []int, k, n int) []int {
	//Misra-Gries Algorithm
	counter := make(map[int]int)
	for i := 0; i < n; i++ {
		x := arr[i]

		count_x, ok := counter[x]
		if !ok {
			if len(counter) < k-1 {
				counter[x] = 1
			} else {
				for key, value := range counter {
					if value == 1 {
						delete(counter, key)
					} else {
						counter[key] = value - 1
					}
				}
			}
		} else {
			counter[x] = count_x + 1
		}
	}

	for key := range counter {
		counter[key] = 0
	}

	for i := 0; i < n; i++ {
		x := arr[i]
		count_x, ok := counter[x]
		if ok {
			counter[x] = count_x + 1
		}
	}

	res := make([]int, 0)

	for key, value := range counter {
		if value > n/k {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	/*
		Design an O(n) algorithm that, given a list of n elements, finds all the
		elements that appear more than n/2 times in the list. Then, design an O(n)
		algorithm that, given a list of n elements, finds all the elements that appear
		more than n/4 times
	*/

	arr := []int{3, 1, 2, 2, 3, 3, 1, 1, 1, 3}
	n := len(arr)
	k := 4
	res := GetElements(arr, k, n)
	fmt.Printf("Elements that appear more than n/%d times are %v\n", k, res)
}
