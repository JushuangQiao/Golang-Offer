package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func createLink(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := Node{val: vals[0], next: nil}
	move := &head
	for _, v := range vals[1:] {
		n := Node{val: v, next: nil}
		move.next = &n
		move = &n
	}
	return &head
}

func printRevertLink(head *Node) {
	if head != nil {
		printRevertLink(head.next)
		fmt.Print(head.val)
	}
}

func main() {
	vals := []int{1, 2, 3, 4}
	h := createLink(vals)
	printRevertLink(h)
}
