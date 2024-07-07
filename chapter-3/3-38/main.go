package main

import "fmt"

type ListNode struct {
	next *ListNode
	key  int
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	key   int
}

type LinkedList struct {
	head *ListNode
}

type Tree struct {
	root *TreeNode
}

func (n *ListNode) String() string {
	s := fmt.Sprintf("key: %d\n", n.key)
	if n.next != nil {
		s += n.next.String()
	}
	return s
}

func (l *LinkedList) String() string {
	if l.head == nil {
		return "Empty List"
	}
	return l.head.String()
}

func insertHelperTree(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{
			key:   key,
			left:  nil,
			right: nil,
		}
	}
	if root.key < key {
		root.left = insertHelperTree(root.left, key)
	} else {
		root.right = insertHelperTree(root.right, key)
	}
	return root
}

func (t *Tree) Insert(k int) {
	t.root = insertHelperTree(t.root, k)
}

func getListNodeFromTreeNode(root *TreeNode) *ListNode {
	if root == nil {
		return nil
	}
	return &ListNode{
		key:  root.key,
		next: nil,
	}
}

func inOrder(root *TreeNode, prev **ListNode) *ListNode {
	if root == nil {
		return nil
	}
	head := inOrder(root.left, prev)
	newNode := getListNodeFromTreeNode(root)
	if *prev != nil {
		(*prev).next = newNode
	} else {
		head = newNode
	}
	*prev = newNode
	inOrder(root.right, prev)
	return head
}

func (t *Tree) GetList() *LinkedList {
	var prev *ListNode
	head := inOrder(t.root, &prev)
	return &LinkedList{head: head}
}

func main() {
	//Write a program to convert a binary search tree into a linked list.
	var arr []int = []int{191, 214, 451, 1, 12, 2, 5, 6, 7, 8, 67, 21}
	var t *Tree = &Tree{}
	for _, i := range arr {
		t.Insert(i)
	}
	var l *LinkedList = t.GetList()
	fmt.Printf("%v", l)
}
