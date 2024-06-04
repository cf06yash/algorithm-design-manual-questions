package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type pair struct {
	idx int
	val rune
}

type stack struct {
	items []*pair
	len   int
}

func (st *stack) push(x rune, idx int) {
	st.items = append(st.items, &pair{idx, x})
	st.len += 1
}

func (st *stack) isEmpty() bool {
	return st.len == 0
}

func (st *stack) pop() (rune, int, error) {
	if st.isEmpty() {
		return -1, -1, errors.New("stack is empty")
	}

	x := st.items[st.len-1]
	st.items = st.items[:st.len-1]
	st.len -= 1
	return x.val, x.idx, nil
}

func (st *stack) peek() (rune, int, error) {
	if st.isEmpty() {
		return -1, -1, errors.New("stack is empty")
	}

	return st.items[st.len-1].val, st.items[st.len-1].idx, nil
}

func (st *stack) lastElem() (rune, int, error) {
	if st.isEmpty() {
		return -1, -1, errors.New("stack is empty")
	}

	return rune(st.items[0].val), st.items[0].idx, nil
}

func hasValidParanthesis(s string) (bool, int) {
	st := stack{}

	for idx, element := range s {
		if rune('(') == element || rune('[') == element || rune('{') == element {
			st.push(element, idx)
		} else if rune(')') == element {
			x, _, err := st.peek()
			if err != nil {
				return false, idx
			}
			if x == rune('(') {
				st.pop()
			} else {
				return false, idx
			}
		} else if rune(']') == element {
			x, _, err := st.peek()
			if err != nil {
				return false, idx
			}
			if x == rune('[') {
				st.pop()
			} else {
				return false, idx
			}
		} else if rune('}') == element {
			x, _, err := st.peek()
			if err != nil {
				return false, idx
			}
			if x == rune('{') {
				st.pop()
			} else {
				return false, idx
			}
		}
	}

	if st.isEmpty() {
		return true, -1
	}

	_, i, _ := st.lastElem()
	return false, i
}

func main() {
	/*
		A common problem for compilers and text editors is determining whether
		the parentheses in a string are balanced and properly nested. For example, the
		string ((())())() contains properly nested pairs of parentheses, while the strings
		)()( and ()) do not. Give an algorithm that returns true if a string contains
		properly nested and balanced parentheses, and false if otherwise. For full credit,
		identify the position of the first offending parenthesis if the string is not properly
		nested and balanced.
	*/
	fmt.Printf("Please Enter Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("Scanned Value: %s\n", s)
	ok, idx := hasValidParanthesis(s)
	if ok {
		fmt.Println("Input String has valid paranthesis")
	} else {
		fmt.Printf("Input does not has valid paranthesis, first error at idx: %d\n", idx)
	}
}
