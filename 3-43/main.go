package main

import "fmt"

type Node struct {
	next *Node
	key  int
}

type LinkedList struct {
	head *Node
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

func (ll *LinkedList) Add(x int) {
	ll.head = addHelper(ll.head, x)
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

func getLastNode(head *Node) *Node {
	if head.next == nil {
		return head
	}
	return getLastNode(head.next)
}

func (ll *LinkedList) AddLoop(loopAfter int) {
	if ll.head == nil {
		return
	}
	if loopAfter < 0 || loopAfter > ll.Size()-1 {
		return
	}

	var last *Node = getLastNode(ll.head)
	var x *Node = ll.head
	for loopAfter > 0 && x != nil {
		loopAfter--
		x = x.next
	}
	fmt.Printf("Creating loop at: %d\n", x.key)
	last.next = x
}

func (ll *LinkedList) FindLoop() *Node {
	if ll.head == nil || ll.head.next == nil {
		return nil
	}
	slow := ll.head
	fast := ll.head

	var found bool = false
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next
		fast = fast.next
		if slow == fast {
			found = true
			break
		}
	}

	if !found {
		return nil
	}

	slow2 := ll.head
	for slow != slow2 {
		slow = slow.next
		slow2 = slow2.next
	}

	return slow
}

func CreateLinkedList(arr []int, loopAfter int) *LinkedList {
	ll := &LinkedList{}
	for i := 0; i < len(arr); i++ {
		ll.Add(arr[i])
	}

	ll.AddLoop(loopAfter)
	return ll
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ll := CreateLinkedList(arr, 6)
	var x *Node = ll.FindLoop()
	if x != nil {
		fmt.Printf("found loop at: %d\n", x.key)
	} else {
		fmt.Printf("no loop\n")
	}

}
