package main

import "fmt"

type Pair struct {
	num   int
	color string
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d,%s) ", p.num, p.color)
}

func GetPair(num int, color string) *Pair {
	return &Pair{
		num:   num,
		color: color,
	}
}

func main() {
	/*
		Assume that we are given n pairs of items as input, where the first item
		is a number and the second item is one of three colors (red, blue, or yellow).
		Further assume that the items are sorted by number. Give an O(n) algorithm
		to sort the items by color (all reds before all blues before all yellows) such that
		the numbers for identical colors stay sorted.
		For example: (1,blue), (3,red), (4,blue), (6,yellow), (9,red) should become (3,red),
		(9,red), (1,blue), (4,blue), (6,yellow).
	*/
	nums := []int{1, 3, 4, 6, 9}
	color := []string{"blue", "red", "blue", "yellow", "red"}

	arr := []*Pair{}
	for i := 0; i < len(nums); i++ {
		arr = append(arr, GetPair(nums[i], color[i]))
	}
	var red []*Pair = []*Pair{}
	var blue []*Pair = []*Pair{}
	var yellow []*Pair = []*Pair{}

	for _, p := range arr {
		if p.color == "red" {
			red = append(red, p)
		} else if p.color == "blue" {
			blue = append(blue, p)
		} else {
			yellow = append(yellow, p)
		}
	}

	arr = append(red, append(blue, yellow...)...)
	fmt.Printf("Sorted by color and number: %v\n", arr)
}
