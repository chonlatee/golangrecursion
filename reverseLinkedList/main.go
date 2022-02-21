package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func reverseListRecursion(n *node) *node { // recursion
	if n == nil || n.next == nil {
		return n
	}

	p := reverseListRecursion(n.next)
	n.next.next = n
	n.next = nil
	return p
}

func reverseList(head *node) *node { // loop
	var prev *node = nil
	var current = head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

func printList(n *node) {
	for n != nil {
		fmt.Printf("%v -> ", n.val)
		n = n.next
	}
}

func main() {

	n60 := &node{
		val:  60,
		next: nil,
	}
	n50 := &node{
		val:  50,
		next: nil,
	}

	n40 := &node{
		val:  40,
		next: nil,
	}

	n30 := &node{
		val:  30,
		next: nil,
	}

	n20 := &node{
		val:  20,
		next: nil,
	}

	n10 := &node{
		val:  10,
		next: nil,
	}

	n10.next = n20
	n20.next = n30
	n30.next = n40
	n40.next = n50
	n50.next = n60

	printList(n10)

	//head := reverseListV1(n10)
	head := reverseListV2(n10)

	fmt.Println()

	fmt.Printf("%+v\n", head)

	printList(head)

}
