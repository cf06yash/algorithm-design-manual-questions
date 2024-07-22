package main

import "fmt"

func searchCol(arr [][]int, li, ri, j, key int) int {
	if li <= ri {
		midi := (li + ri) / 2
		if arr[midi][j] == key {
			return midi
		} else if arr[midi][j] > key {
			return searchCol(arr, li, midi-1, j, key)
		} else {
			return searchCol(arr, midi+1, ri, j, key)
		}
	}
	return -1
}

func searchRow(arr [][]int, lj, rj, key int) (int, int) {
	if lj <= rj {
		midj := (lj + rj) / 2
		if arr[0][midj] == key {
			return 0, midj
		} else if arr[0][midj] > key {
			return searchRow(arr, lj, midj-1, key)
		} else {
			i := searchCol(arr, 0, len(arr)-1, midj, key)
			if i != -1 {
				return i, midj
			} else {
				return searchRow(arr, midj+1, rj, key)
			}
		}
	}

	return -1, -1
}

func main() {
	M := [][]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
		{3, 5, 7, 9, 11},
		{4, 6, 8, 10, 12},
		{5, 7, 9, 11, 13},
	}

	// Function to print the matrix
	printMatrix := func(matrix [][]int) {
		for _, row := range matrix {
			fmt.Println(row)
		}
	}

	fmt.Println("Matrix M:")
	printMatrix(M)

	testCases := []int{7, 1, 13, 0, 14, 6, 15}

	for _, x := range testCases {
		fmt.Printf("\nSearching for %d\n", x)
		resulti, resultj := searchRow(M, 0, len(M[0])-1, x)
		if resulti != -1 {
			fmt.Printf("Result: %v\n", M[resulti][resultj])
		} else {
			fmt.Println("Not found")
		}
	}
}
