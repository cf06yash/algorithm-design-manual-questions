package main

import "fmt"

func getSum(arr []int, l, r int) int {
	if l >= r {
		return arr[l]
	} else {
		mid := (l + r) / 2
		rightMax := arr[mid+1]
		leftMax := arr[mid]

		left := 0
		for i := mid; i >= l; i-- {
			left += arr[i]
			if left > leftMax {
				leftMax = left
			}
		}

		right := 0
		for i := mid + 1; i <= r; i++ {
			right += arr[i]
			if right > rightMax {
				rightMax = right
			}
		}

		css := leftMax + rightMax

		lss := getSum(arr, l, mid-1)
		rss := getSum(arr, mid+1, r)

		return max(css, max(lss, rss))
	}
}

func main() {
	//The largest subrange problem, discussed in Section 5.6, takes an array A of
	//n numbers, and asks for the index pair i and j that maximizes S = jk=i A[k].
	//Give an O(n) algorithm for largest subrange.

	arr := []int{-17, 5, 3, -10, 6, 1, 4, -3, 8, 1, -13, 4}
	sum := getSum(arr, 0, len(arr)-1)
	fmt.Printf("For array %v max subarray sum is %d\n", arr, sum)
}
