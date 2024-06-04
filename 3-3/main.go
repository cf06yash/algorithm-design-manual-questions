package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	val  int
	next *node
}

type list struct {
	head *node
}

func (l *list) add(x int) {
	if l.head == nil {
		l.head = &node{
			val:  x,
			next: nil,
		}
		return
	}

	var n *node = l.head
	for n.next != nil {
		n = n.next
	}
	n.next = &node{
		val:  x,
		next: nil,
	}
}

func reverseHelper(l *list, head *node, pred *node) {
	if head == nil {
		l.head = pred
		return
	}
	reverseHelper(l, head.next, head)
	head.next = pred
}

func (l *list) reverseList() {
	if l.head == nil {
		return
	}
	reverseHelper(l, l.head, nil)
}

func (l *list) print() {
	if l.head == nil {
		return
	}
	var n *node = l.head
	for n != nil {
		fmt.Printf("%d ", n.val)
		n = n.next
	}
	fmt.Printf("\n")
}

func main() {
	/*
		Give an algorithm to reverse the direction of a given singly linked list. In
		other words, after the reversal all pointers should now point backwards. Your
		algorithm should take linear time.
	*/

	fmt.Printf("Please Enter Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("Scanned Value: %s\n", s)
	a := strings.Split(s, ",")

	var l *list = &list{}
	for _, x := range a {
		y, err := strconv.Atoi(strings.TrimSpace(x))
		if err == nil {
			l.add(y)
		}
	}

	l.reverseList()
	l.print()
}
