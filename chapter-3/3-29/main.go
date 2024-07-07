package main

import "fmt"

type FenwickTree struct {
	tree []int
}

type Pair struct {
	key   int
	value int
}

func getNext(i int) int {
	x := ^i + 1
	y := i & x
	return i + y
}

func (t *FenwickTree) build(arr []*Pair, idx int, treeIdx int) {
	//fmt.Printf("Values: arr: %v, idx %d, treeIdx %d\n", arr, idx, treeIdx)
	if treeIdx > len(t.tree)-1 {
		return
	}
	t.tree[treeIdx] += arr[idx].value
	t.build(arr, idx, getNext(treeIdx))
}

func (t *FenwickTree) updateHelper(treeIdx int, value int) {
	if treeIdx > len(t.tree)-1 {
		return
	}
	t.tree[treeIdx] += value
	t.updateHelper(getNext(treeIdx), value)
}

func (t *FenwickTree) Update(arr *[]*Pair, key, value int) {
	idx := searchKey(arr, key, 0, len(*arr)-1)
	if idx == -1 {
		return
	}
	(*arr)[idx].value += value
	t.updateHelper(idx+1, value)
}

func (t *FenwickTree) Insert(arr *[]*Pair, key, value int) {
	idx := searchIdxToInsert(arr, key, 0, len(*arr)-1)
	newArr := make([]*Pair, len(*arr)+1)
	for i := 0; i < idx; i++ {
		newArr[i] = (*arr)[i]
	}
	newArr[idx] = &Pair{
		key:   key,
		value: value,
	}
	for i := idx + 1; i < len(newArr); i++ {
		newArr[i] = (*arr)[i-1]
	}

	*arr = newArr
	t.tree = make([]int, len(*arr)+1)
	for i := 0; i < len(*arr); i++ {
		t.build(*arr, i, i+1)
	}
}

func searchIdxToInsert(arr *[]*Pair, key, l, r int) int {
	if l > r {
		return l
	}
	mid := (l + r) / 2
	if (*arr)[mid].key < key {
		return searchIdxToInsert(arr, key, mid+1, r)
	} else {
		return searchIdxToInsert(arr, key, l, mid-1)
	}
}

func searchKey(arr *[]*Pair, key, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if (*arr)[mid].key == key {
		return mid
	} else if (*arr)[mid].key > key {
		return searchKey(arr, key, l, mid-1)
	} else {
		return searchKey(arr, key, mid+1, r)
	}
}

func (t *FenwickTree) Delete(arr *[]*Pair, key int) {
	idx := searchKey(arr, key, 0, len(*arr))
	if idx == -1 {
		return
	}
	newArr := make([]*Pair, len(*arr)-1)
	copy(newArr[0:idx], (*arr)[0:idx])
	copy(newArr[idx:], (*arr)[idx+1:])
	*arr = newArr
	for i := idx; i < len(t.tree); i++ {
		t.tree[idx+1] = 0
	}
	t.tree = make([]int, len(*arr)+1)
	for i := 0; i < len(*arr); i++ {
		t.build(*arr, i, i+1)
	}
}

func CreateNewTree(arr []*Pair) *FenwickTree {
	t := &FenwickTree{
		tree: make([]int, len(arr)+1),
	}
	for i := 0; i < len(arr); i++ {
		t.build(arr, i, i+1)
	}
	return t
}

func main() {
	/*
		Extend the data structure of the previous problem to support insertions and
		deletions. Each element now has both a key and a value. An element is accessed
		by its key, but the addition operation is applied to the values. The Partial sum
		operation is different.
		• Add(k,y) – Add the value y to the item with key k.
		• Insert(k,y) – Insert a new item with key k and value y.
		• Delete(k) – Delete the item with key k.
		• Partial-sum(k) – Return the sum of all the elements currently in the set
		whose key is less than k, that is,
		i<k xi.
		The worst-case running time should still be O(n log n) for any sequence of O(n)
		operations
	*/
	arr := []*Pair{{key: 1, value: 1}, {key: 2, value: 2}, {key: 4, value: 4}, {key: 5, value: 5}, {key: 6, value: 6}}
	t := CreateNewTree(arr)

	fmt.Println("Initial Array and Fenwick Tree:")
	for _, pair := range arr {
		fmt.Printf("Key: %d, Value: %d\n", pair.key, pair.value)
	}
	fmt.Println("Fenwick Tree:", t.tree)

	t.Insert(&arr, 3, 3)
	fmt.Println("\nArray and Fenwick Tree after Inserting (3, 3):")
	for _, pair := range arr {
		fmt.Printf("Key: %d, Value: %d\n", pair.key, pair.value)
	}
	fmt.Println("Fenwick Tree:", t.tree)

	t.Update(&arr, 4, 1)
	fmt.Println("\nArray and Fenwick Tree after Updating (4, 1):")
	for _, pair := range arr {
		fmt.Printf("Key: %d, Value: %d\n", pair.key, pair.value)
	}
	fmt.Println("Fenwick Tree:", t.tree)

	t.Delete(&arr, 2)
	fmt.Println("\nArray and Fenwick Tree after Deleting key 2:")
	for _, pair := range arr {
		fmt.Printf("Key: %d, Value: %d\n", pair.key, pair.value)
	}
	fmt.Println("Fenwick Tree:", t.tree)
}
