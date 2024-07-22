package main

import "fmt"

func Examine(A []rune, i int) (rune, bool) {
	if i >= len(A) {
		return '0', false
	}
	return A[i], true
}

func Swap(A *[]rune, i, j int) {
	(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
}

func CustomSort(arr *[]rune) {
	if len(*arr) == 0 {
		return
	}
	countR := 0
	countW := 0
	countB := 0

	for i := 0; i < len(*arr); i++ {
		ball, ok := Examine(*arr, i)
		if !ok {
			break
		}
		if ball == 'R' {
			countR++
		} else if ball == 'W' {
			countW++
		} else if ball == 'B' {
			countB++
		}
	}

	Swapper(arr, 0, len(*arr)-1, countR, 'R')
	Swapper(arr, countR, len(*arr)-1, countW+countR, 'W')
}

func Swapper(arr *[]rune, i, j, count int, color rune) {
	for ; i < count; i++ {
		ball, ok := Examine(*arr, i)
		if !ok {
			break
		}
		if ball != color {
			ballFromRight, ok := Examine(*arr, j)
			for ; j > i && ok && ballFromRight != color; ballFromRight, ok = Examine(*arr, j) {
				j--
			}
			Swap(arr, i, j)
		}
	}
}

func main() {
	/*
		Suppose an array A consists of n elements, each of which is red, white, or
		blue. We seek to sort the elements so that all the reds come before all the whites,
		which come before all the blues. The only operations permitted on the keys are:
		• Examine(A,i) – report the color of the ith element of A.
		• Swap(A,i,j) – swap the ith element of A with the jth element.
		Find a correct and efficient algorithm for red–white–blue sorting. There is a
		linear-time solution.
	*/

	A := []rune{'R', 'W', 'B', 'R', 'W', 'B', 'B', 'R', 'W'}
	fmt.Printf("%c\n", A)
	CustomSort(&A)
	fmt.Printf("%c\n", A)
}
