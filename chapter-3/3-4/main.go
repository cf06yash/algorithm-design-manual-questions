package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items  []int
	minIdx []int
}

func (s *Stack) isEmpty() bool {
	return len(s.items) <= 0
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
	min, err := s.findmin()
	if err != nil || min >= i {
		s.minIdx = append(s.minIdx, len(s.items)-1)
	}
}

func (s *Stack) findmin() (int, error) {
	if len(s.minIdx) == 0 {
		return -99999, errors.New("stack is empty")
	}
	return s.items[s.minIdx[len(s.minIdx)-1]], nil
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return -99999, errors.New("stack is empty")
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	if s.minIdx[len(s.minIdx)-1] == idx {
		s.minIdx = s.minIdx[:len(s.minIdx)-1]
	}
	return item, nil
}

func main() {
	/*
		Design a stack S that supports S.push(x), S.pop(), and S.findmin(), which
		returns the minimum element of S. All operations should run in constant time
	*/
	fmt.Printf("Creating Stack\n")
	st := &Stack{
		items:  make([]int, 0),
		minIdx: make([]int, 0),
	}
	fmt.Printf("Allowed Queries push, pop, findmin, quit\n")
	for {
		fmt.Printf("Please enter your query: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		query := strings.Split(s, " ")
		if query[0] == "push" {

			i, err := strconv.Atoi(query[1])
			if err != nil {
				fmt.Printf("Could not scan input %v\n", err)
				continue
			}
			st.Push(i)
		} else if query[0] == "pop" {
			i, err := st.Pop()
			if err != nil {
				fmt.Printf("Error could not pop item %v\n", err)
			} else {
				fmt.Printf("Popped item: %d\n", i)
			}
		} else if query[0] == "findmin" {
			i, err := st.findmin()
			if err != nil {
				fmt.Printf("Error could find min item %v\n", err)
			} else {
				fmt.Printf("Min item: %d\n", i)
			}
		} else if query[0] == "quit" {
			fmt.Printf("Quit command initiated, exiting flow\n")
			break
		} else {
			fmt.Printf("Unexpected input, user entered: %s, allowed values: push,pop,findmin,quit\n", s)
		}
	}

}
