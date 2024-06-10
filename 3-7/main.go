package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val  int
	next *node
}

type List struct {
	head *node
	tail *node
}

func NewList() *List {
	l := &List{
		head: nil,
		tail: &node{val: -11, next: nil},
	}
	return l
}

func (l *List) Add(i int) *node {
	if l.head == nil {
		l.head = &node{
			val:  i,
			next: l.tail,
		}
		return l.head
	}
	n := &node{
		val:  i,
		next: nil,
	}

	c := l.head
	for c.next != l.tail {
		c = c.next
	}
	c.next = n
	n.next = l.tail
	return n
}

func (l *List) PrintList() {
	if l.head == nil {
		return
	}
	c := l.head
	for c != nil {
		fmt.Printf("%d ", c.val)
		c = c.next
	}
	fmt.Println("")
}

func (l *List) DeleteNode(x *node) {
	if l.head == nil {
		return
	}
	if x == nil {
		return
	}
	nextNode := x.next
	if nextNode == nil {
		fmt.Printf("error next node should not be nil")
		return
	}
	x.val = nextNode.val
	x.next = nextNode.next
	if nextNode == l.tail {
		l.tail = x
	}
}

func main() {
	/*
		Work out the details of supporting constant-time deletion from a singly
		linked list as per the footnote from page 79, ideally to an actual implementation.
		Support the other operations as efficiently as possible.
	*/
	arr := []int{1, 2, 3, 4, 5}
	shuffle(arr)
	fmt.Printf("Shuffled array: %v\n", arr)
	l := NewList()

	// I am taking pointer of node with val 3 to delete, and have shuffled the array so that 3 is in a different position each time
	var nodeToDel *node = nil
	for i := 0; i < len(arr); i++ {
		y := l.Add(arr[i])
		if arr[i] == 3 {
			nodeToDel = y
		}
	}
	l.PrintList()
	fmt.Println("Deleting 3 from the list")
	l.DeleteNode(nodeToDel)
	l.PrintList()
}

func shuffle(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
