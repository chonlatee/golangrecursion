package main

import "fmt"

type node struct {
	val  int
	next *node
}

func createFirstLinkedList() *node {
	n1 := &node{
		val:  1,
		next: nil,
	}

	n8 := &node{
		val:  8,
		next: nil,
	}

	n15 := &node{
		val:  15,
		next: nil,
	}

	n20 := &node{
		val:  20,
		next: nil,
	}

	n30 := &node{
		val:  30,
		next: nil,
	}

	n1.next = n8
	n8.next = n15
	n15.next = n20
	n20.next = n30

	return n1
}

func createSecondLinkedList() *node {
	n3 := &node{
		val:  3,
		next: nil,
	}

	n7 := &node{
		val:  7,
		next: nil,
	}

	n14 := &node{
		val:  14,
		next: nil,
	}

	n21 := &node{
		val:  21,
		next: nil,
	}

	n28 := &node{
		val:  28,
		next: nil,
	}

	n3.next = n7
	n7.next = n14
	n14.next = n21
	n21.next = n28

	return n3
}

func mergeTwoSortedLinkedListRecursion(first *node, second *node) *node { // recursion
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	if first.val < second.val {
		first.next = mergeTwoSortedLinkedListRecursion(first.next, second)
		return first
	} else {
		second.next = mergeTwoSortedLinkedListRecursion(first, second.next)
		return second
	}
}

func mergeTwoSortedLinkedList(first *node, second *node) *node { // loop
	var t *node
	var cur *node
	var head *node

	head = second
	if first.val < second.val {
		head = first
	}

	cur = head

	for first != nil && second != nil {
		if first.val < second.val {
			t = first
			first = first.next
		} else {
			t = second
			second = second.next
		}

		cur.next = t
		cur = cur.next

	}

	for second != nil {
		cur.next = second
		cur = cur.next
		second = second.next
	}

	for first != nil {
		cur.next = first
		cur = cur.next
		first = first.next
	}

	return head

}

func printList(n *node) {
	for n != nil {
		fmt.Printf("%v -> ", n.val)
		n = n.next
	}
}

func main() {

	f := createFirstLinkedList()
	s := createSecondLinkedList()

	fmt.Println()
	printList(f)
	fmt.Println()
	printList(s)
	fmt.Println()

	n := mergeTwoSortedLinkedListRecursion(f, s)
	printList(n)

	fmt.Println()

	f = createFirstLinkedList()
	s = createSecondLinkedList()

	fmt.Println()
	printList(f)
	fmt.Println()
	printList(s)
	fmt.Println()

	n = mergeTwoSortedLinkedList(f, s)
	printList(n)

}
