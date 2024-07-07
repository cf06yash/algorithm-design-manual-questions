package main

import (
	"errors"
	"fmt"
)

type Store struct {
	a     []int
	b     []int
	count int
}

func (s *Store) Search(x int) int {
	if x > len(s.a) || x <= 0 {
		return -1
	}
	if s.a[x-1] < len(s.b) && s.a[x-1] >= 0 && s.b[s.a[x-1]] == x {
		return s.a[x-1]
	}
	return -1
}

func (s *Store) Insert(x int) {
	if x > len(s.a) || x <= 0 {
		fmt.Printf("invalid x\n")
		return
	}
	if s.count >= len(s.b) {
		fmt.Printf("Store is full\n")
		return
	}
	if s.Search(x) == -1 {
		s.a[x-1] = s.count
		s.b[s.count] = x
		s.count++
	}
}

func (s *Store) Delete(x int) {
	idx := s.Search(x)
	if idx == -1 {
		return
	}

	//update latest element in a with idx
	y := s.b[s.count-1]

	s.a[y-1] = idx
	s.b[idx] = y

	s.a[x-1] = -1
	s.count--
}

func InitializeStore(n, m int) (*Store, error) {
	if n < m {
		return nil, errors.New("n should be greater than or equal to m")
	}
	return &Store{
		a:     make([]int, n),
		b:     make([]int, m),
		count: 0,
	}, nil
}

func main() {
	/*
		Design a data structure that allows one to search, insert, and delete an
		integer X in O(1) time (i.e., constant time, independent of the total number of
		integers stored). Assume that 1 ≤ X ≤ n and that there are m + n units of
		space available, where m is the maximum number of integers that can be in the
		table at any one time. (Hint: use two arrays A[1..n] and B[1..m].) You are not
		allowed to initialize either A or B, as that would take O(m) or O(n) operations.
		This means the arrays are full of random garbage to begin with, so you must be
		very careful.
	*/
	var n = 10
	var m = 5
	store, err := InitializeStore(n, m)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	store.Insert(1)
	store.Insert(2)
	store.Insert(3)
	store.Insert(5)
	store.Delete(1)
	fmt.Printf("A: %v\n", store.a)
	fmt.Printf("B: %v\n", store.b)
	fmt.Printf("count: %d\n", store.count)
}
