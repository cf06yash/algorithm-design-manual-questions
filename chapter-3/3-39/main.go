package main

import "fmt"

type Node struct {
	next *Node
	key  int
}

type LinkedList struct {
	head *Node
}

func (n *Node) String() string {
	s := fmt.Sprintf("key: %d ", n.key)
	if n.next != nil {
		s += n.next.String()
	}
	return s
}

func (l *LinkedList) String() string {
	if l.head == nil {
		return "Empty List\n"
	}
	return l.head.String()
}

func addHelper(head *Node, x int) *Node {
	if head == nil {
		return &Node{
			key:  x,
			next: nil,
		}
	}
	head.next = addHelper(head.next, x)
	return head
}

func reverseRecursive(head *Node, prev *Node, newHead **Node) {
	if head == nil {
		return
	}
	reverseRecursive(head.next, head, newHead)
	if head.next == nil {
		*newHead = head
	}
	head.next = prev
}

func (ll *LinkedList) ReverseRecur() {
	if ll.head == nil {
		return
	}
	var newHead *Node
	reverseRecursive(ll.head, nil, &newHead)
	ll.head = newHead
}

func (ll *LinkedList) Add(x int) {
	ll.head = addHelper(ll.head, x)
}

func (ll *LinkedList) Reverse() {
	if ll.head == nil {
		return
	}
	var prevprev *Node
	prev := ll.head
	head := prev.next

	for head != nil {
		prev.next = prevprev
		prevprev = prev
		prev = head
		head = head.next
	}
	prev.next = prevprev
	ll.head = prev
}

func CreateList(arr []int) *LinkedList {
	ll := &LinkedList{}
	for i := 0; i < len(arr); i++ {
		ll.Add(arr[i])
	}
	return ll
}

func main() {
	//Implement an algorithm to reverse a linked list. Now do it without recursion.
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ll := CreateList(arr)
	fmt.Printf("Before: %v\n", ll)
	ll.ReverseRecur()
	fmt.Printf("After: %v\n", ll)
	ll.Reverse()
	fmt.Printf("Reverse Again: %v\n", ll)
}
