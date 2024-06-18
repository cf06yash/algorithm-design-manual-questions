package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

type ListNode struct {
	prev *ListNode
	next *ListNode
	val  int
}

func (n ListNode) String() string {
	if n.prev != nil && n.next != nil {
		return fmt.Sprintf("Val: %d, Prev: %v, Next: %v\n", n.val, n.prev.val, n.next.val)
	} else if n.prev != nil && n.next == nil {
		return fmt.Sprintf("Val: %d, Prev: %v, Next: nil\n", n.val, n.prev.val)
	} else if n.next != nil && n.prev == nil {
		return fmt.Sprintf("Val: %d, Prev: nil, Next: %v\n", n.val, n.next.val)
	} else {
		return fmt.Sprintf("Val: %d, Prev: nil, Next: nil\n", n.val)
	}

}

type DoubleLinkedList struct {
	head *ListNode
}

func addNode(root, n *Node) *Node {
	if root == nil {
		return n
	}
	if root.val >= n.val {
		root.left = addNode(root.left, n)
	} else {
		root.right = addNode(root.right, n)
	}

	return root
}

func (t *BST) Add(val int) {
	t.root = addNode(t.root, &Node{
		val:   val,
		left:  nil,
		right: nil,
	})
}

func inOrder(root *Node, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.left, result)
	*result = append(*result, root.val)
	inOrder(root.right, result)
}

func (t *BST) InOrder() []int {
	res := make([]int, 0)
	inOrder(t.root, &res)
	return res
}

func Add(head, n *ListNode) {
	if head == nil {
		return
	}
	c := head
	for c.next != nil {
		c = c.next
	}
	c.next = n
	n.prev = c
}

func (ll *DoubleLinkedList) Add(val int) {
	if ll.head == nil {
		ll.head = &ListNode{
			prev: nil,
			next: nil,
			val:  val,
		}
		return
	}
	Add(ll.head, &ListNode{
		prev: nil,
		next: nil,
		val:  val,
	})
}

func (ll *DoubleLinkedList) merge(a, b []int) {
	if len(a) == 0 && len(b) == 0 {
		return
	}
	if len(a) == 0 {
		for _, v := range b {
			ll.Add(v)
		}
		return
	}
	if len(b) == 0 {
		for _, v := range a {
			ll.Add(v)
		}
		return
	}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		for i < len(a) && a[i] < b[j] {
			ll.Add(a[i])
			i++
		}
		if i >= len(a) {
			break
		}
		for j < len(b) && b[j] <= a[i] {
			ll.Add(b[j])
			j++
		}
	}

	for i < len(a) {
		ll.Add(a[i])
		i++
	}

	for j < len(b) {
		ll.Add(b[j])
		j++
	}
}

func (ll *DoubleLinkedList) Traverse() []*ListNode {
	result := make([]*ListNode, 0)
	c := ll.head
	for c != nil {
		result = append(result, c)
		c = c.next
	}
	return result
}

func main() {
	a := &BST{}
	b := &BST{}
	for i := 0; i < 30; i++ {
		randomNumber := rand.Intn(2) + 1
		if randomNumber == 1 {
			a.Add(i)
		} else {
			b.Add(i)
		}
	}
	resA := a.InOrder()
	resB := b.InOrder()
	fmt.Printf("Tree A: %v\n", resA)
	fmt.Printf("Tree B: %v\n", resB)
	var ll *DoubleLinkedList = &DoubleLinkedList{}
	ll.merge(resA, resB)
	fmt.Printf("Merged list:\n %v\n", ll.Traverse())
}
