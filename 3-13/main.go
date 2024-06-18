package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

type Pair struct {
	val    int
	parent *Node
	linkNo int
}

func (p Pair) String() string {
	return fmt.Sprintf("%d", p.val)
}

func addNode(root, n *Node) *Node {
	if root == nil {
		if n.val == 6 { //todo introducing wrong placement
			n.val = 5
		} else if n.val == 5 {
			n.val = 6
		}
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

func inOrder(root *Node, result *[]*Pair, parent *Node, link int) {
	if root == nil {
		return
	}
	inOrder(root.left, result, root, 1)
	*result = append(*result, &Pair{
		val:    root.val,
		parent: parent,
		linkNo: link,
	})
	inOrder(root.right, result, root, 2)
}

func (t *BST) InOrder() []*Pair {
	var res []*Pair = make([]*Pair, 0)
	inOrder(t.root, &res, nil, 0)
	return res
}

func getWrongNodes(res []*Pair) (*Pair, *Pair) {
	if len(res) == 0 {
		return nil, nil
	}
	var a *Pair
	for i := 0; i < len(res)-1; i++ {
		if res[i].val > res[i+1].val {
			a = res[i]
			break
		}
	}
	var b *Pair
	for i := len(res) - 1; i > 0; i-- {
		if res[i].val < res[i-1].val {
			b = res[i]
			break
		}
	}

	return a, b
}

func swap(a, b *Pair) {
	ap := a.parent
	bp := b.parent
	var childa *Node
	var childb *Node
	if a.linkNo == 1 && b.linkNo == 1 {
		childa = ap.left
		childb = bp.left
		ap.left = childb
		bp.left = childa
	} else if a.linkNo == 1 {
		childa = ap.left
		childb = bp.right
		ap.left = childb
		bp.right = childa
	} else if a.linkNo == 2 && b.linkNo == 2 {
		childa = ap.right
		childb = bp.right
		ap.right = childb
		bp.right = childa
	} else {
		childa = ap.right
		childb = bp.left
		ap.right = childb
		bp.left = childa
	}
	var temp *Node
	temp = childa.left
	childa.left = childb.left
	childb.left = temp
	temp = childa.right
	childa.right = childb.right
	childb.right = temp

}

func (t *BST) Correction() {
	res := t.InOrder()
	a, b := getWrongNodes(res)
	if a == nil || b == nil {
		return
	}
	fmt.Printf("Wrong nodes: A - %v B - %v\n", a, b)
	swap(a, b)
}

func main() {
	t := &BST{}
	arr := []int{3, 4, 6, 7, 9, 1, 2, 8, 5, 0}
	for i := 0; i < 10; i++ {
		t.Add(arr[i])
	}
	res := t.InOrder()
	fmt.Printf("%v\n", res)
	t.Correction()
	res = t.InOrder()
	fmt.Printf("%v\n", res)
}
