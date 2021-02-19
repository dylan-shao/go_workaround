package main

import (
"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func helper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	
	helper(t.Left, ch)
	ch <- t.Value
	
	helper(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	helper(t, ch)
	
	close(ch)
}

func compare(ch1 chan int, ch2 chan int) bool{
	v1,ok1 := <- ch1
	v2,ok2 := <- ch2
	if !ok1 && !ok2 {
		return true
	}
	if v1 != v2 {
		return false
	}
	return compare(ch1, ch2)
}
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	return compare(ch1, ch2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

