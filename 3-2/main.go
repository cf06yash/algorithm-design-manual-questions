package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	items []int
}

func (st *stack) push(i int) {
	st.items = append(st.items, i)
}

func (st *stack) isEmpty() bool {
	return len(st.items) == 0
}

func (st *stack) pop() int {
	if st.isEmpty() {
		return -1
	}
	x := st.items[len(st.items)-1]
	st.items = st.items[:len(st.items)-1]
	return x
}

func (st *stack) peek() int {
	if st.isEmpty() {
		return -1
	}
	return st.items[len(st.items)-1]
}

func findLongestValidParanthesis(s string) int {
	st := stack{}
	var count int = 0
	for _, element := range s {
		if element == rune('(') {
			st.push(int(element))
		} else if element == rune(')') {
			x := st.peek()
			if x == int(rune('(')) {
				count += 1
				st.pop()
			}
		}
	}

	return 2 * count
}

func main() {

	/*
		Give an algorithm that takes a string S consisting of opening and closing
		parentheses, say )()(())()()))())))(, and finds the length of the longest balanced
		parentheses in S, which is 12 in the example above. (Hint: The solution is not
		necessarily a contiguous run of parenthesis from S.)
	*/

	fmt.Printf("Please Enter Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("Scanned Value: %s\n", s)
	fmt.Printf("Longest Valid Paranthesis : %d\n", findLongestValidParanthesis(s))
}
