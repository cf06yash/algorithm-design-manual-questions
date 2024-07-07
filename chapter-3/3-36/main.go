package main

import "fmt"

type Node struct {
	next *Node
	key  int
}

func (n *Node) String() string {
	s := fmt.Sprintf("key: %d\n", n.key)
	if n.next != nil {
		s += n.next.String()
	}
	return s
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) String() string {
	if l.head == nil {
		return "Empty List"
	}
	return l.head.String()
}

func add(head, n *Node) *Node {
	if head == nil {
		return n
	}
	head.next = add(head.next, n)
	return head
}

func (l *LinkedList) Add(x int) {
	n := &Node{
		next: nil,
		key:  x,
	}
	l.head = add(l.head, n)
}

func sizeHelper(count *int, head *Node) {
	if head == nil {
		return
	}
	(*count)++
	sizeHelper(count, head.next)
}

func (l *LinkedList) Size() int {
	count := 0
	sizeHelper(&count, l.head)
	return count
}

func getNodeHelper(mid int, head *Node) *Node {
	if head == nil {
		return nil
	}
	if mid == 0 {
		return head
	}
	return getNodeHelper(mid-1, head.next)
}

func (l *LinkedList) GetNode(mid int) *Node {
	return getNodeHelper(mid, l.head)
}

func (l *LinkedList) Middle() *Node {
	size := l.Size()
	mid := (size - 1) / 2
	return l.GetNode(mid)
}

func CreateList(arr []int) *LinkedList {
	l := &LinkedList{
		head: nil,
	}
	for i := 0; i < len(arr); i++ {
		l.Add(arr[i])
	}
	return l
}

func main() {
	//Write a function to find the middle node of a singly linked list.
	arr := []int{5, 1, 12, 567, 1, 122, 699, 9, 10}
	ll := CreateList(arr)
	fmt.Printf("%v\n", ll)
	var m *Node = ll.Middle()
	if m == nil {
		fmt.Printf("Could not find middle :(\n")
	} else {
		fmt.Printf("Middle element: %v\n", m.key)
	}
}
