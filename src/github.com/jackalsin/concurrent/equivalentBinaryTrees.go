package main

import (
	"fmt"
)
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

func walkHelper(t *Tree, ch chan int) {
	if t!= nil {
		walkHelper(t.Left, ch)
		ch <- t.Value
		walkHelper(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	//t1 := Tree.New(1)
	//t2 := Tree.New(1)
	//fmt.Println(Same(t1, t2))
}

func Traversal(t *Tree) {
	if t.Left != nil {
		Traversal(t.Left)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		Traversal(t.Right)
	}
}
